package services

import (
	"encoding/json"
	"flexagent/models"
	"flexagent/pkg/config"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/flexlet/utils"
	"github.com/google/uuid"
	"github.com/robfig/cron"
)

var CronJobService *cronjobServiceImpl

type cronJobScheduler struct {
	crontab *cron.Cron
	cronjob *models.CronJob
}

type cronjobServiceImpl struct {
	mutex      *sync.Mutex
	schedulers map[string]*cronJobScheduler
}

func NewCronJobScheduler(cronjob *models.CronJob) (*cronJobScheduler, error) {
	cronJobName := cronjob.Spec.Name
	// schedule cron job
	crontab := cron.New()

	// function to call on schedule
	errCron := crontab.AddFunc(cronjob.Spec.Schedule, func() {
		// submit job when scheduled
		job, errJob := JobService.Submit(cronjob.Spec.Jobspec)
		if errJob != nil {
			utils.LogPrintf(utils.LOG_ERROR, "services.CronJobScheduler",
				"cron job '%s' schedule failed, reason: %s", cronJobName, errJob.Error())
			return
		}

		// record history jobs
		cronjob.Jobs = append(cronjob.Jobs, job.Urn)

		// delete oldest jobs when exceed max history number
		for len(cronjob.Jobs) > config.AgentConfig.CronJobMaxHistoryJobs {
			oldJobUrn := cronjob.Jobs[0]
			cronjob.Jobs = cronjob.Jobs[1:]
			errDel := JobService.Delete(oldJobUrn)
			if errDel != nil {
				utils.LogPrintf(utils.LOG_ERROR, "services.CronJobScheduler",
					"cron job '%s' clean history job '%s' failed, reason: %s", cronJobName, oldJobUrn, errDel.Error())
			}
		}
	})
	if errCron != nil {
		return nil, errCron
	}

	// start scheduler
	scheduler := &cronJobScheduler{cronjob: cronjob, crontab: crontab}
	scheduler.start()

	return scheduler, nil
}

func (s *cronJobScheduler) stop() {
	s.crontab.Stop()
	s.cronjob.Status = models.CronJobStatusStoped
}

func (s *cronJobScheduler) start() {
	s.crontab.Start()
	s.cronjob.Status = models.CronJobStatusRunning
}

func NewCronJobService() *cronjobServiceImpl {
	schedulers := make(map[string]*cronJobScheduler, config.AgentConfig.MaxCronJobs)

	// list cron job files
	cronjobFiles, errList := ioutil.ReadDir(config.AgentConfig.CronJobsPath)
	if errList != nil {
		utils.LogPrintf(utils.LOG_ERROR, "services.NewCronJobService",
			"list cron jobs failed, reason: %s", errList.Error())
		os.Exit(0)
	}

	if len(cronjobFiles) > config.AgentConfig.MaxCronJobs {
		utils.LogPrintf(utils.LOG_ERROR, "services.NewCronJobService",
			"Too many cronjobs, maximum: %d", config.AgentConfig.MaxCronJobs)
		os.Exit(0)
	}

	for i := 0; i < len(cronjobFiles); i++ {
		cronJobName := cronjobFiles[i].Name()
		cronjobFile := config.AgentConfig.CronJobsPath + "/" + cronJobName

		// load cron job
		cronjob, errLoad := readCronJobFile(cronjobFile)
		if errLoad != nil {
			utils.LogPrintf(utils.LOG_ERROR, "services.NewCronJobService",
				"load cron job '%s' failed, reason: %s", cronJobName, errLoad.Error())
			os.Exit(0)
		}

		// create scheduler
		scheduler, errCron := NewCronJobScheduler(cronjob)
		if errCron != nil {
			utils.LogPrintf(utils.LOG_ERROR, "services.NewCronJobService",
				"create cron job '%s' failed, reason: %s", cronJobName, errCron.Error())
			os.Exit(0)
		}

		// add to the map
		schedulers[cronJobName] = scheduler
	}

	// initialize schedulers
	return &cronjobServiceImpl{
		mutex:      &sync.Mutex{},
		schedulers: schedulers,
	}
}

func (s *cronjobServiceImpl) BatchSubmit(specs []*models.CronJobSpec) ([]*models.CronJob, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var cronjobs []*models.CronJob

	// check maximum number of cornjobs
	if len(s.schedulers)+len(specs) > config.AgentConfig.MaxCronJobs {
		return cronjobs, fmt.Errorf("Too many cronjobs, maximum: %d", config.AgentConfig.MaxCronJobs)
	}

	// submit cronjobs
	for i := 0; i < len(specs); i++ {
		cronjob := &models.CronJob{
			ID:     uuid.NewString(),
			Status: models.CronJobStatusStoped,
			Spec:   specs[i],
			Jobs:   make([]string, 0),
		}

		// write to file
		cronjobFile := config.AgentConfig.CronJobsPath + "/" + cronjob.ID
		if err := writeCronJobFile(cronjob, cronjobFile); err != nil {
			return cronjobs, fmt.Errorf("dump cronjob '%s' failed, reason: %s", cronjob.Spec.Name, err.Error())
		}

		// create scheduler
		scheduler, err := NewCronJobScheduler(cronjob)
		if err != nil {
			return cronjobs, fmt.Errorf("schedule cronjob '%s' failed, reason: %s", cronjob.Spec.Name, err.Error())
		}

		// add to the map
		s.schedulers[cronjob.ID] = scheduler

		cronjobs = append(cronjobs, cronjob)
	}

	return cronjobs, nil
}

func (s *cronjobServiceImpl) List(name *string) ([]*models.CronJob, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var cronjobs []*models.CronJob
	for _, scheduler := range s.schedulers {
		cronjobName := scheduler.cronjob.Spec.Name
		if name != nil && !strings.Contains(cronjobName, *name) {
			continue
		}
		cronjobs = append(cronjobs, scheduler.cronjob)
	}
	return cronjobs, nil
}

func (s *cronjobServiceImpl) Query(id string) (*models.CronJob, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check exist
	scheduler, exist := s.schedulers[id]
	if !exist {
		return nil, fmt.Errorf("cronjob '%s' does not exist", id)
	}

	return scheduler.cronjob, nil
}

func (s *cronjobServiceImpl) Update(id string, spec *models.CronJobSpec) (*models.CronJob, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check exist
	scheduler, exist := s.schedulers[id]
	if !exist {
		return nil, fmt.Errorf("cronjob '%s' does not exist", id)
	}

	// stop old scheduler
	scheduler.stop()

	// replace job spec (keep old history jobs)
	newCronjob := scheduler.cronjob
	newCronjob.Spec = spec

	// write to file
	cronjobFile := config.AgentConfig.CronJobsPath + "/" + id
	if err := writeCronJobFile(newCronjob, cronjobFile); err != nil {
		return nil, fmt.Errorf("dump cronjob '%s' failed, reason: %s", id, err.Error())
	}

	// create new scheduler
	newScheduler, err := NewCronJobScheduler(newCronjob)
	if err != nil {
		return nil, fmt.Errorf("update cronjob '%s' failed, reason: ", id, err.Error())
	}

	// add to map
	s.schedulers[id] = newScheduler

	return newCronjob, nil
}

func (s *cronjobServiceImpl) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check exist
	scheduler, exist := s.schedulers[id]
	if !exist {
		return fmt.Errorf("cronjob '%s' does not exist", id)
	}

	// stop scheduler
	scheduler.stop()

	// delete history jobs
	for i := 0; i < len(scheduler.cronjob.Jobs); i++ {
		if err := JobService.Delete(scheduler.cronjob.Jobs[i]); err != nil {
			// ignore job delete fail
			continue
		}
	}

	// delete dump file
	cronjobFile := config.AgentConfig.CronJobsPath + "/" + id
	if err := os.Remove(cronjobFile); err != nil {
		return fmt.Errorf("cronjob '%s' delete failed, reason: %s", id, err.Error())
	}

	// delete from map
	delete(s.schedulers, id)

	return nil
}

func (s *cronjobServiceImpl) Stop(id string) (*models.CronJob, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check exist
	scheduler, exist := s.schedulers[id]
	if !exist {
		return nil, fmt.Errorf("cronjob '%s' does not exist", id)
	}

	// stop scheduler
	scheduler.stop()

	return scheduler.cronjob, nil
}

func (s *cronjobServiceImpl) Start(id string) (*models.CronJob, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check exist
	scheduler, exist := s.schedulers[id]
	if !exist {
		return nil, fmt.Errorf("cronjob '%s' does not exist", id)
	}

	// stop scheduler
	scheduler.start()

	return scheduler.cronjob, nil
}

// dump job to file
func writeCronJobFile(cronjob *models.CronJob, file string) error {
	var (
		plainText  []byte
		cipherText []byte
		err        error
	)

	if plainText, err = json.Marshal(*cronjob); err != nil {
		return err
	}

	if cipherText, err = CryptoService.Crypto.Encrypt(plainText); err != nil {
		return err
	}

	return ioutil.WriteFile(file, cipherText, utils.MODE_PERM_RW)
}

// read file to job
func readCronJobFile(file string) (*models.CronJob, error) {
	var (
		plainText  []byte
		cipherText []byte
		err        error
	)

	if cipherText, err = ioutil.ReadFile(file); err != nil {
		return nil, err
	}

	if plainText, err = CryptoService.Crypto.Decrypt(cipherText); err != nil {
		return nil, err
	}

	var cronjob models.CronJob

	if err = json.Unmarshal(plainText, &cronjob); err != nil {
		return nil, err
	}

	return &cronjob, nil
}

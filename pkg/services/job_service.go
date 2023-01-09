package services

import (
	"bufio"
	"encoding/json"
	"flexagent/models"
	"flexagent/pkg/config"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"gitee.com/yaohuiwang/utils"
	"github.com/google/uuid"
)

var JobService *jobServiceImpl

const (
	JOB_OUT_FILE       = "job.out"
	JOB_DMP_FILE       = "job.dmp"
	DefaultOutputLines = 10
)

type jobServiceImpl struct {
	jobThreadPool *utils.ThreadPool
	jobQueue      map[string]*utils.Script
	mutex         *sync.Mutex
}

func NewJobService() *jobServiceImpl {
	// initialize thread pool
	return &jobServiceImpl{
		jobThreadPool: utils.NewThreadPool(config.AgentConfig.JobQueueSize, config.AgentConfig.JobThreadPoolSize),
		jobQueue:      make(map[string]*utils.Script),
		mutex:         &sync.Mutex{},
	}
}

func (s *jobServiceImpl) GracefulShutdown() {
	// wait all jobs until end
	for _, script := range s.jobQueue {
		jobPath := filepath.Dir(script.Spec.Out)
		script.Wait()
		s.queryAndUpdate(jobPath, false, nil, nil)
	}
}

func (s *jobServiceImpl) BatchSubmit(specs []*models.JobSpec, wait *bool) ([]*models.Job, error) {
	var jobs []*models.Job
	for i := 0; i < len(specs); i++ {
		job, err := s.Submit(specs[i])
		if err != nil {
			return jobs, err
		}
		jobs = append(jobs, job)
	}

	// wait jobs finish
	if wait != nil && *wait {
		for i := 0; i < len(jobs); i++ {
			if newJob, err := s.Wait(jobs[i].Urn); err != nil {
				utils.LogPrintf(utils.LOG_ERROR, "jobServiceImpl.BatchSubmit", "wait job failed: %s", err.Error())
			} else {
				jobs[i] = newJob
			}
		}
	}

	return jobs, nil
}

func (s *jobServiceImpl) Submit(spec *models.JobSpec) (*models.Job, error) {
	jobId := uuid.NewString()
	jobUrn := fmt.Sprintf("%s:jobs:%s:%s", spec.Plugin, spec.Operation, jobId)
	jobPath := fmt.Sprintf("%s/%s", config.AgentConfig.PluginsLogPath, strings.ReplaceAll(jobUrn, ":", "/"))
	utils.MkdirIfNotExist(jobPath)

	pluginConfig, pluginExist := config.PluginConfigs[spec.Plugin]

	if !pluginExist {
		return nil, fmt.Errorf("plugin does not exist")
	}

	scriptFile, scriptExist := pluginConfig.Operations[spec.Operation]
	if !scriptExist {
		return nil, fmt.Errorf("plugin has no such operation")
	}

	workDir := config.AgentConfig.PluginsPath + "/" + spec.Plugin
	scriptPath := workDir + "/" + scriptFile
	scriptSpec := &utils.ScriptSpec{
		Path:        scriptPath,
		Dir:         workDir,
		Args:        spec.Args,
		Env:         spec.Env,
		SysProcAttr: &syscall.SysProcAttr{Setpgid: true},
		Out:         jobPath + "/" + JOB_OUT_FILE,
	}

	script, err := utils.NewScript(scriptSpec, func() {
		// job done callback
		s.queryAndUpdate(jobPath, false, nil, nil)
	})

	if err != nil {
		return nil, err
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := s.jobThreadPool.Put(script); err != nil {
		return nil, err
	}
	s.jobQueue[jobId] = script

	job := &models.Job{
		ID:   jobId,
		Urn:  jobUrn,
		Spec: spec,
		Status: &models.JobStatus{
			State: models.JobStatusStateWaiting,
		},
	}

	jobFile := fmt.Sprintf("%s/%s", jobPath, JOB_DMP_FILE)

	if err := writeJobFile(job, jobFile); err != nil {
		return nil, err
	}

	return job, nil
}

// dump job to file
func writeJobFile(job *models.Job, file string) error {
	var (
		plainText  []byte
		cipherText []byte
		err        error
	)

	if plainText, err = json.Marshal(*job); err != nil {
		return err
	}

	if cipherText, err = CryptoService.Crypto.Encrypt(plainText); err != nil {
		return err
	}

	return ioutil.WriteFile(file, cipherText, utils.MODE_PERM_RW)
}

// read file to job
func readJobFile(file string) (*models.Job, error) {
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

	var job models.Job

	if err = json.Unmarshal(plainText, &job); err != nil {
		return nil, err
	}

	return &job, nil
}

func (s *jobServiceImpl) List(plugin *string, operation *string, startTimeBegin *int64, startTimeEnd *int64) ([]*models.Job, error) {
	var jobs []*models.Job
	err := filepath.Walk(config.AgentConfig.PluginsLogPath, func(path string, info os.FileInfo, err1 error) error {
		if info.Name() != JOB_DMP_FILE {
			return nil
		}
		tokens := strings.Split(path, "/")
		len := len(tokens)
		jobOperation := tokens[len-3]
		jobPlugin := tokens[len-5]

		// only for linux
		jobStartTime := info.Sys().(*syscall.Stat_t).Ctim.Sec

		if plugin != nil && jobPlugin != *plugin {
			return nil
		}
		if operation != nil && jobOperation != *operation {
			return nil
		}
		if startTimeBegin != nil && jobStartTime < *startTimeBegin {
			return nil
		}
		if startTimeEnd != nil && jobStartTime > *startTimeEnd {
			return nil
		}

		job, err := s.queryAndUpdate(filepath.Dir(path), false, nil, nil)
		if err != nil {
			return err
		}

		jobs = append(jobs, job)

		return nil
	})

	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func parseJobUrn(jobUrn string) (string, string) {
	jobPath := config.AgentConfig.PluginsLogPath + "/" + strings.ReplaceAll(jobUrn, ":", "/")
	jobId := filepath.Base(jobPath)
	return jobId, jobPath
}

func (s *jobServiceImpl) Query(jobUrn string, outputLineStart *int32, outputLineLimit *int32) (*models.Job, error) {
	_, jobPath := parseJobUrn(jobUrn)

	if !utils.FileExist(jobPath) {
		return nil, fmt.Errorf("job does not exist")
	}

	// queue job
	if job, err := s.queryAndUpdate(jobPath, true, outputLineStart, outputLineLimit); err != nil {
		return nil, err
	} else {
		return job, nil
	}
}

// read job output file
func readJobOutFile(jobOutFile string, outputLineStart *int32, outputLineLimit *int32) (*models.JobOutput, error) {
	file, err := os.Open(jobOutFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		lastLine int32 = 0
		moreLine bool  = true
		limit    int   = DefaultOutputLines
		lines    []string
	)

	if outputLineLimit != nil {
		limit = int(*outputLineLimit)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		line = strings.TrimRight(line, "\n")
		lastLine++

		// start line number is set
		if outputLineStart != nil && lastLine >= *outputLineStart {
			lines = append(lines, line)
			if len(lines) == limit {
				break
			}
		}

		// start line number is not set, get last lines
		if outputLineStart == nil {
			lines = append(lines, line)
			if len(lines) > limit {
				lines = lines[1:]
			}
		}

		if err == io.EOF {
			moreLine = false
			break
		}
	}

	jobOutput := &models.JobOutput{
		LastLine: lastLine,
		MoreLine: moreLine,
		Lines:    lines,
	}
	return jobOutput, nil

}

// query job, update job file if still running
func (s *jobServiceImpl) queryAndUpdate(jobPath string, withOutput bool, outputLineStart *int32, outputLineLimit *int32) (*models.Job, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	jobId := filepath.Base(jobPath)
	jobFile := jobPath + "/" + JOB_DMP_FILE
	jobOutFile := jobPath + "/" + JOB_OUT_FILE

	job, err := readJobFile(jobFile)
	if err != nil {
		return nil, fmt.Errorf("read job filed, reason: %s", err.Error())
	}

	// update if job is still on the queue
	if script, exist := s.jobQueue[jobId]; exist {
		// update status
		job.Status.State = script.Status

		// update exit code
		if script.Cmd.ProcessState != nil {
			exitCode := int32(script.Cmd.ProcessState.ExitCode())
			job.Status.ExitCode = &exitCode
		}

		// update job file
		writeJobFile(job, jobFile)

		// delete from job queue
		if script.Status != utils.ScriptStatusWaiting && script.Status != utils.ScriptStatusRunning {
			delete(s.jobQueue, jobId)
		}
	}

	// get output
	if withOutput {
		jobOutput, err := readJobOutFile(jobOutFile, outputLineStart, outputLineLimit)
		if err != nil {
			return nil, fmt.Errorf("read job output filed, reason: %s", err.Error())
		}

		job.Status.Output = jobOutput
	}
	return job, nil
}

func (s *jobServiceImpl) Delete(jobUrn string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	jobId, jobPath := parseJobUrn(jobUrn)
	if !utils.FileExist(jobPath) {
		return fmt.Errorf("job does not exist")
	}

	// if exist on job queue
	if script, exist := s.jobQueue[jobId]; exist {
		if script.Cmd.ProcessState != nil && script.Cmd.ProcessState.Exited() {
			// if exited, delete from job queue
			delete(s.jobQueue, jobId)
		} else {
			// if not start or still running
			return fmt.Errorf("job is waiting or running, please wait")
		}
	}

	return os.RemoveAll(jobPath)
}

func (s *jobServiceImpl) Input(jobUrn string, input *models.JobInput) (*models.Job, error) {
	jobId, jobPath := parseJobUrn(jobUrn)

	if !utils.FileExist(jobPath) {
		return nil, fmt.Errorf("job does not exist")
	}

	script, exist := s.jobQueue[jobId]
	if !exist || script.Cmd.Process == nil {
		return nil, fmt.Errorf("job is not in running state")
	}

	// expect before input
	if input.Expect != nil && input.Timeout != nil {
		matched := script.Expect(*input.Expect, time.Duration(*input.Timeout)*time.Second)
		if !matched {
			return nil, fmt.Errorf("expect timeout")
		}
	}

	// input data
	if err := script.Input(input.Data); err != nil {
		return nil, fmt.Errorf("write to script stdin failed, reason: %s", err.Error())
	}

	// queue job
	if job, err := s.queryAndUpdate(jobPath, true, nil, nil); err != nil {
		return nil, err
	} else {
		return job, nil
	}
}

func (s *jobServiceImpl) Kill(jobUrn string, force *bool) (*models.Job, error) {
	jobId, jobPath := parseJobUrn(jobUrn)

	if !utils.FileExist(jobPath) {
		return nil, fmt.Errorf("job does not exist")
	}

	script, exist := s.jobQueue[jobId]
	if !exist || script.Cmd.Process == nil {
		return nil, fmt.Errorf("job is not in running state")
	}

	// kill job
	if force != nil && *force {
		syscall.Kill(-script.Cmd.Process.Pid, syscall.SIGKILL)
	} else {
		syscall.Kill(-script.Cmd.Process.Pid, syscall.SIGINT)
	}

	// wait until exited
	script.Wait()

	// queue job
	if job, err := s.queryAndUpdate(jobPath, true, nil, nil); err != nil {
		return nil, err
	} else {
		return job, nil
	}
}

func (s *jobServiceImpl) Wait(jobUrn string) (*models.Job, error) {
	jobId, jobPath := parseJobUrn(jobUrn)

	if !utils.FileExist(jobPath) {
		return nil, fmt.Errorf("job does not exist")
	}

	script, exist := s.jobQueue[jobId]
	if exist && (script.Status == utils.ScriptStatusWaiting || script.Status == utils.ScriptStatusRunning) {
		// wait until exited
		script.Wait()
	}

	// queue job
	if job, err := s.queryAndUpdate(jobPath, true, nil, nil); err != nil {
		return nil, err
	} else {
		return job, nil
	}
}

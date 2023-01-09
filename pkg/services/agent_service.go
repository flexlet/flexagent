package services

import (
	"flexagent/models"
	"flexagent/pkg/config"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/flexlet/utils"
)

var AgentService *agentServiceImpl

type agentServiceImpl struct {
	mutex       *sync.Mutex
	readyStatus string
}

func NewAgentService() *agentServiceImpl {
	return &agentServiceImpl{
		mutex:       &sync.Mutex{},
		readyStatus: models.ReadyStatusStatusReady,
	}
}

func (s *agentServiceImpl) SetReady() {
	s.readyStatus = models.ReadyStatusStatusReady
}

func (s *agentServiceImpl) SetNotReady() {
	s.readyStatus = models.ReadyStatusStatusNotReady
}

func (s *agentServiceImpl) Readyz() *models.ReadyStatus {
	return &models.ReadyStatus{
		Status: s.readyStatus,
	}
}

func (s *agentServiceImpl) Healthz() *models.HealthStatus {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	probes := make(map[string]map[string]models.HealthProbe)
	overallStatus := config.STATUS_UNKNOWN

	for pluginName, pluginConfig := range config.PluginConfigs {
		probes[pluginName] = make(map[string]models.HealthProbe)
		for probeName, probeScript := range pluginConfig.Probes {
			workDir := config.AgentConfig.PluginsPath + "/" + pluginName
			logDir := fmt.Sprintf("%s/%s/probes", config.AgentConfig.PluginsLogPath, pluginName)
			if err := utils.MkdirIfNotExist(logDir); err != nil {
				continue
			}
			status, output := execProbe(probeName, probeScript, workDir, logDir)
			probes[pluginName][probeName] = models.HealthProbe{
				Name:    probeName,
				Message: output,
				Status:  config.ProbeStatusEnum[status],
			}
			if status > overallStatus {
				overallStatus = status
			}
		}
	}

	return &models.HealthStatus{
		Probes: probes,
		Status: config.ProbeStatusEnum[overallStatus],
	}
}

func execProbe(probeName string, probeScript string, workDir string, logDir string) (int, *string) {
	scriptSpec := &utils.ScriptSpec{
		Path: workDir + "/" + probeScript,
		Dir:  workDir,
		Out:  logDir + "/" + probeName,
		Env:  config.ProbeStatusEnv,
	}

	script, err := utils.NewScript(scriptSpec, nil)
	if err != nil {
		return config.STATUS_UNKNOWN, nil
	}
	script.Run()
	status := script.Cmd.ProcessState.ExitCode()
	output, err := ioutil.ReadFile(script.Spec.Out)
	if err != nil {
		return status, nil
	}

	message := strings.TrimRight(string(output), "\n")
	return status, &message
}

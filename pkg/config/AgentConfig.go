package config

import (
	"io/ioutil"
	"os"

	"github.com/flexlet/utils"
	"gopkg.in/yaml.v3"
)

var AgentConfig = struct {
	LogLevel              uint8             `yaml:"log_level"`
	Keystores             map[string]string `yaml:"keystores"`
	Algorithum            string            `yaml:"algorithum"`
	VaultsPath            string            `yaml:"vaults_path"`
	CronJobsPath          string            `yaml:"cronjobs_path"`
	PluginsPath           string            `yaml:"plugins_path"`
	PluginsLogPath        string            `yaml:"plugins_log_path"`
	Plugins               []string          `yaml:"plugins"`
	JobQueueSize          int               `yaml:"job_queue_size"`
	JobThreadPoolSize     int               `yaml:"job_thread_pool_size"`
	MaxCronJobs           int               `yaml:"max_cronjobs"`
	CronJobMaxHistoryJobs int               `yaml:"cronjob_max_history_jobs"`
}{}

var PluginConfigs map[string]*PluginConfig = make(map[string]*PluginConfig)

func LoadAgentConfig(conf string) {
	// load config file
	data, err := ioutil.ReadFile(conf)
	if err != nil {
		utils.LogPrintf(utils.LOG_ERROR, "config.LoadAgentConfig",
			"load '%s' failed, reason: %s", conf, err.Error())
		os.Exit(0)
	}
	err = yaml.Unmarshal(data, &AgentConfig)
	if err != nil {
		utils.LogPrintf(utils.LOG_ERROR, "config.LoadAgentConfig",
			"load '%s' failed, reason: %s", conf, err.Error())
		os.Exit(0)
	}

	// create vault dir
	utils.MkdirIfNotExist(AgentConfig.VaultsPath)

	// create cronjob dir
	utils.MkdirIfNotExist(AgentConfig.CronJobsPath)

	// create plugin log dir
	utils.MkdirIfNotExist(AgentConfig.PluginsLogPath)

	// load plugins
	for i := 0; i < len(AgentConfig.Plugins); i++ {
		pluginName := AgentConfig.Plugins[i]
		pluginConfigFile := AgentConfig.PluginsPath + "/" + pluginName + "/plugin.yaml"
		pluginConfig, err := LoadPluginConfig(pluginConfigFile)
		if err != nil {
			utils.LogPrintf(utils.LOG_ERROR, "config.LoadAgentConfig",
				"Load plugin '%s' failed, reason: %s", pluginName, err.Error())
			os.Exit(0)
		}
		PluginConfigs[pluginName] = pluginConfig
	}

}

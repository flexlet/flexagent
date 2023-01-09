package config

import (
	"fmt"
	"io/ioutil"

	"gitee.com/yaohuiwang/utils"
	"gopkg.in/yaml.v3"
)

type PluginConfig struct {
	Name       string            `yaml:"name"`
	Probes     map[string]string `yaml:"probes"`
	Operations map[string]string `yaml:"operations"`
}

func LoadPluginConfig(conf string) (*PluginConfig, error) {
	// load config file
	data, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, err
	}
	var pluginConfig PluginConfig
	err = yaml.Unmarshal(data, &pluginConfig)
	if err != nil {
		return nil, err
	}

	// validate probe scripts
	for key, script := range pluginConfig.Probes {
		if !utils.FileExist(AgentConfig.PluginsPath + "/" + pluginConfig.Name + "/" + script) {
			return nil, fmt.Errorf("load probe '%s' error, '%s' does not exist", key, script)
		}
	}

	// validate operations scripts
	for key, script := range pluginConfig.Operations {
		if !utils.FileExist(AgentConfig.PluginsPath + "/" + pluginConfig.Name + "/" + script) {
			return nil, fmt.Errorf("load operation '%s' error, '%s' does not exist", key, script)
		}
	}

	return &pluginConfig, nil
}

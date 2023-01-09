package config

import (
	"flexagent/models"
	"fmt"
)

const (
	PROJECT string = "Node Agent"
	VERSION string = "0.1.0"
)

const (
	STATUS_UNKNOWN int = iota
	STATUS_HEALTHY
	STATUS_WARNING
	STATUS_CRITICAL
)

var (
	ProbeStatusEnv = []string{
		fmt.Sprintf("%s=%d", "STATUS_UNKNOWN", STATUS_UNKNOWN),
		fmt.Sprintf("%s=%d", "STATUS_HEALTHY", STATUS_HEALTHY),
		fmt.Sprintf("%s=%d", "STATUS_WARNING", STATUS_WARNING),
		fmt.Sprintf("%s=%d", "STATUS_CRITICAL", STATUS_CRITICAL),
	}
	ProbeStatusEnum []string = []string{
		models.HealthProbeStatusUnknown,
		models.HealthProbeStatusHealthy,
		models.HealthProbeStatusWarning,
		models.HealthProbeStatusCritical,
	}
)

# An example of app plugin

A app plugin must contains the following directives:

```yaml

# application name
name: <app_name>

# service health probe directives
probes:

  # service status check script
  # script exit code should be: ${STATUS_CRITICAL} ${STATUS_WARNING} ${STATUS_HEALTY}
  service_status: "probes/service_status.sh"

# service lifecycle management operations
operations:
  # service init script
  service_init: "operations/service_init.sh"

  # service start script
  service_start: "operations/service_start.sh"
  
  # service stop script
  service_stop: "operations/service_stop.sh"
  
  # service restart script
  service_restart: "operations/service_restart.sh"
  
  # optional, print logs script, should have params: -s SINCE -u UNTIL
  service_logs: "operations/service_logs.sh"

```
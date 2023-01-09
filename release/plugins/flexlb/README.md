# flexlb plugin

## Probes

### Probe service_status

```sh
/opt/flexagent/plugins/flexlb/probes/service_status.sh 

flexlb: [active]
```

## Operations

Test on local:

```sh
# test on local
SERVER="http://127.0.0.1:18080"
alias JSON="python -c \"import json; import sys; print(json.dumps(json.loads(sys.stdin.read()), indent=4, ensure_ascii=False));\""
```

### Operation service_cfg

Configure service params

```sh
# configure members
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
          "plugin": "flexlb",
          "operation": "service_cfg",
          "args": ["-m","192.168.1.193:8001"],
          "env": []
         }]'

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}" | JSON

```


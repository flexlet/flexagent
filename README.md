# Flexible Agent

A flexible agent framework

## Build & install

Build:

```sh
 # edit build config to select plugins (under release/plugins)
vi build/profile

# generate certs
sh build/gencerts.sh

# build
sh build/build.sh

```

Install:

```sh
# copy to node
scp release/target/flexagent-amd64-v0.1.0.tar.gz root@192.168.1.193:/root/

# install on node
tar -zxf flexagent-amd64-v0.1.0.tar.gz
cd flexagent-amd64-v0.1.0

sh install.sh

# check service
systemctl status flexagent

```

Test on local:

```sh
# test on local
SERVER="http://127.0.0.1:18080"
alias JSON="python -c \"import json; import sys; print(json.dumps(json.loads(sys.stdin.read()), indent=4, ensure_ascii=False));\""
```

## Agent health

Agent ready status and health probes.

```sh

# ready status
curl ${SERVER}/api/v1/readyz

{"status":"ready"}

# health status
curl ${SERVER}/api/v1/healthz | JSON

{
  "probes": {
    "linux": {
      "cpu_usage": {
        "message": "cpu usage: 0.96%",
        "name": "cpu_usage",
        "status": "healthy"
      },
      "fs_usage": {
        "message": "/ => Healthy (3.42%), /boot => Healthy (14.11%), /home => Healthy (0.12%), ",
        "name": "fs_usage",
        "status": "healthy"
      },
      "mem_usage": {
        "message": "mem usage: 3.42%",
        "name": "mem_usage",
        "status": "healthy"
      }
    }
  },
  "status": "healthy"
}

```

## Crypto

Crypto for encrypt or decrypt

### Encrypt/Decrypt

Encrypt or decrypt k/v data

```sh
# encrypt k/v raw plain text
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/crypto/encrypt" \
     -d '{"format":"raw","data":{"key1":"data1","key2":"data2"}}'

# encrypt k/v base64 plain text
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/crypto/encrypt" \
     -d '{"format":"base64","data":{"key1":"ZGF0YTE=","key2":"ZGF0YTI="}}'

# decrypt k/v raw cipher text
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/crypto/decrypt" \
     -d '{
        "format":"raw",
        "data":{
          "key1":"AAAAAgAAAAAAAAABAAAAAAAAAAlX6tkg+bmJ193v7HmHUBUmRKDe1gMu+D5Da9tKAAAAAAEAAAEAAAAAAAAAFfecYXkTd/5axapbNI0Ypec8i6Tfgg==",
          "key2":"AAAAAgAAAAAAAAABAAAAAAAAAAlX6tkg+bmJ17sGecf9v0eUVJsOwudkWbK6pYrGAAAAAAEAAAEAAAAAAAAAFbyuvOf/CMna7bvxCF8Ykv2jozuTwA=="
        }
      }'

# decrypt k/v base64 cipher text
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/crypto/decrypt" \
     -d '{
        "format":"base64",
        "data":{
          "key1":"QUFBQUFnQUFBQUFBQUFBQkFBQUFBQUFBQUFsWDZ0a2crYm1KMTBrck05UWJ4OEw2YWh6b2JQM2wrNHA1WHEwVUFBQUFBQUVBQUFFQUFBQUFBQUFBRmF0U0RCVUxHNWV1NDJDdzV0cWhIMU5LY2krS1dRPT0=",
          "key2":"QUFBQUFnQUFBQUFBQUFBQkFBQUFBQUFBQUFsWDZ0a2crYm1KMTVuOFVrajA5WXJoRkNxZ1dBd0tIait3S1F2cUFBQUFBQUVBQUFFQUFBQUFBQUFBRlRhcDloNzVWK3FwRzdqb0ROUk1tVWZ3S1dISXpnPT0="
        }
      }'
```

### Secret encrypt / decrypt

Encrypt or decrypt k8s secret

```sh
# encrypt k8 secret
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/crypto/secret/encrypt" \
     -d '{
        "apiVersion": "v1",
        "kind": "Secret",
        "type": "Opaque",
        "metadata": {"name": "secret1"},
        "data": {"key1":"ZGF0YTE=","key2":"ZGF0YTI="}
      }'

# decrypt k8s secret
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/crypto/secret/decrypt" \
     -d '{
        "apiVersion": "v1",
        "kind": "Secret",
        "type": "Opaque",
        "metadata": {"name": "secret1"},
        "data": {
          "key1":"QUFBQUFnQUFBQUFBQUFBQkFBQUFBQUFBQUFsWDZ0a2crYm1KMTMxODErbnhaRnA2L1ZWNEJYa3FZVUFDUmtGdEFBQUFBQUVBQUFFQUFBQUFBQUFBRmVFUGVSZG1JSFg1T3ZEWWFFcGtLT09rVHpDaHpBPT0=",
          "key2":"QUFBQUFnQUFBQUFBQUFBQkFBQUFBQUFBQUFsWDZ0a2crYm1KMTNtUWl5U2lKRDFjbGc3ZllWYWhHckZsZFFMZ0FBQUFBQUVBQUFFQUFBQUFBQUFBRmRUSFdwWi9OeWgyS0hKalhoWk5HNDFGbG1sWGdBPT0="
        }
      }'
```

### Vault

Vault to store k/v data

```sh
# create vault
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/crypto/vault/vault1" \
     -d '{"key1":"data1","key2":"data2"}'

# update vault
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X PUT "${SERVER}/api/v1/crypto/vault/vault1" \
     -d '{"key1":"update1","key3":"data3"}'

# query vault keys
curl "${SERVER}/api/v1/crypto/vault/vault1?keys=key1,key2"

# query vault
curl "${SERVER}/api/v1/crypto/vault/vault1"

# delete vault keys
curl -X DELETE "${SERVER}/api/v1/crypto/vault/vault1?keys=key1,key2"

# delete vault
curl -X DELETE "${SERVER}/api/v1/crypto/vault/vault1"

# list vaults
curl "${SERVER}/api/v1/crypto/vaults"
```

## Run pugin jobs

[Linux plugin](release/plugins/linux/README.md)

## Run cronjobs

```sh
# create cronjobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/cronjobs" \
     -d '[{
       "name": "ping_local_1m",
       "schedule": "0 */1 * * * *",
       "jobspec": {"plugin":"linux","operation":"ping","args":["-a","127.0.0.1","-c","1"]}
      }]'

[
  {
    "id": "f6c97cd2-0978-4230-95c0-35128cc1fdb8",
    "jobs": [],
    "spec": {
      "jobspec": {
        "args": [
          "-a",
          "127.0.0.1",
          "-c",
          "1"
        ],
        "env": null,
        "operation": "ping",
        "plugin": "linux"
      },
      "name": "ping_local_1m",
      "schedule": "0 */1 * * * *"
    },
    "status": "running"
  }
]

CRONJOBID="f6c97cd2-0978-4230-95c0-35128cc1fdb8"

# list cronjobs
curl "${SERVER}/api/v1/cronjobs?name=ping"

[
  {
    "id": "f6c97cd2-0978-4230-95c0-35128cc1fdb8",
    "jobs": [
      "linux:jobs:ping:c10a1180-b29e-4700-9200-85829da301bd"
    ],
    "spec": {
      "jobspec": {
        "args": [
          "-a",
          "127.0.0.1",
          "-c",
          "1"
        ],
        "env": null,
        "operation": "ping",
        "plugin": "linux"
      },
      "name": "ping_local_1m",
      "schedule": "0 */1 * * * *"
    },
    "status": "running"
  }
]

# query cronjob
curl "${SERVER}/api/v1/cronjobs/${CRONJOBID}"

# update cronjob
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X PUT "${SERVER}/api/v1/cronjobs/${CRONJOBID}" \
     -d '{
       "name": "ping_local_5m",
       "schedule": "0 */5 * * * *",
       "jobspec": {"plugin":"linux","operation":"ping","args":["-a","127.0.0.1","-c","1"]}
      }'

# stop cronjob
curl -X POST "${SERVER}/api/v1/cronjobs/${CRONJOBID}/stop"

# start cronjob
curl -X POST "${SERVER}/api/v1/cronjobs/${CRONJOBID}/start"

# delete cronjob
curl -X DELETE "${SERVER}/api/v1/cronjobs/${CRONJOBID}"

```
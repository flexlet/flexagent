# General linux plugin

## Probes

### Probe cpu_usage

```sh
/opt/flexagent/plugins/linux/probes/cpu_usage.sh 

cpu usage: 0.92%
```

### Probe mem_usage

```sh
/opt/flexagent/plugins/linux/probes/mem_usage.sh 

mem usage: 4.46%
```

### Probe fs_usage

```sh
/opt/flexagent/plugins/linux/probes/fs_usage.sh 

max fs usage: 14.11%, / => Healthy (3.42%), /boot => Healthy (14.11%), /home => Healthy (0.12%)
```

## Operations

Test on local:

```sh
# test on local
SERVER="http://127.0.0.1:18080"
alias JSON="python -c \"import json; import sys; print(json.dumps(json.loads(sys.stdin.read()), indent=4, ensure_ascii=False));\""
```

### Operation ifcfg

Config ip address for a network interface.

```sh
# submit jobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{"plugin":"linux","operation":"ifcfg","args":["-d","enp4s3","-a","192.168.2.193","-m","255.255.255.0"]}]'

[
  {
    "id": "8dbedb7e-1ea7-4407-b399-46cecf1bdec8",
    "spec": {
      "args": [
        "-d",
        "enp4s3",
        "-a",
        "192.168.2.193",
        "-m",
        "255.255.255.0"
      ],
      "env": null,
      "operation": "ifcfg",
      "plugin": "linux"
    },
    "status": {
      "state": "waiting"
    },
    "urn": "linux:jobs:ifcfg:8dbedb7e-1ea7-4407-b399-46cecf1bdec8"
  }
]

JOBURN="linux:jobs:ifcfg:8dbedb7e-1ea7-4407-b399-46cecf1bdec8"

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}"

{
  "id": "8dbedb7e-1ea7-4407-b399-46cecf1bdec8",
  "spec": {
    "args": [
      "-d",
      "enp4s3",
      "-a",
      "192.168.2.193",
      "-m",
      "255.255.255.0"
    ],
    "env": null,
    "operation": "ifcfg",
    "plugin": "linux"
  },
  "status": {
    "exitCode": 0,
    "output": {
      "lastLine": 9,
      "lines": [
        "Connection 'enp4s3' successfully deactivated (D-Bus active path: /org/freedesktop/NetworkManager/ActiveConnection/3)",
        "Connection successfully activated (D-Bus active path: /org/freedesktop/NetworkManager/ActiveConnection/4)",
        "4: enp4s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000",
        "    link/ether 28:6e:d4:88:f7:b3 brd ff:ff:ff:ff:ff:ff",
        "    inet 192.168.2.193/24 brd 192.168.2.255 scope global noprefixroute enp4s3",
        "       valid_lft forever preferred_lft forever",
        "    inet6 fe80::2a6e:d4ff:fe88:f7b3/64 scope link tentative ",
        "       valid_lft forever preferred_lft forever",
        ""
      ],
      "moreLine": false
    },
    "state": "exited"
  },
  "urn": "linux:jobs:ifcfg:8dbedb7e-1ea7-4407-b399-46cecf1bdec8"
}

# check job output
cat ${LOGDIR}/linux/jobs/ifcfg/8dbedb7e-1ea7-4407-b399-46cecf1bdec8/job.out

# check job dump file
cat ${LOGDIR}/linux/jobs/ifcfg/8dbedb7e-1ea7-4407-b399-46cecf1bdec8/job.dmp | /usr/local/bin/crypto-tool -decrypt

# delete job
curl -X DELETE "${SERVER}/api/v1/jobs/${JOBURN}"

```

### Operation ping

Ping a host.

```sh
# submit job
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{"plugin":"linux","operation":"ping","args":["-a","127.0.0.1","-c","100"]}]'

[
  {
    "id": "f072dba4-5c93-4136-876d-cf6b8f6cab9e",
    "spec": {
      "args": [
        "-a",
        "127.0.0.1",
        "-c",
        "100"
      ],
      "env": null,
      "operation": "ping",
      "plugin": "linux"
    },
    "status": {
      "state": "waiting"
    },
    "urn": "linux:jobs:ping:f072dba4-5c93-4136-876d-cf6b8f6cab9e"
  }
]

JOBURN="linux:jobs:ping:f072dba4-5c93-4136-876d-cf6b8f6cab9e"

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}"

{
  "id": "f072dba4-5c93-4136-876d-cf6b8f6cab9e",
  "spec": {
    "args": [
      "-a",
      "127.0.0.1",
      "-c",
      "100"
    ],
    "env": null,
    "operation": "ping",
    "plugin": "linux"
  },
  "status": {
    "output": {
      "lastLine": 76,
      "lines": [
        "64 bytes from 127.0.0.1: icmp_seq=66 ttl=64 time=0.026 ms",
        "64 bytes from 127.0.0.1: icmp_seq=67 ttl=64 time=0.021 ms",
        "64 bytes from 127.0.0.1: icmp_seq=68 ttl=64 time=0.024 ms",
        "64 bytes from 127.0.0.1: icmp_seq=69 ttl=64 time=0.014 ms",
        "64 bytes from 127.0.0.1: icmp_seq=70 ttl=64 time=0.021 ms",
        "64 bytes from 127.0.0.1: icmp_seq=71 ttl=64 time=0.023 ms",
        "64 bytes from 127.0.0.1: icmp_seq=72 ttl=64 time=0.019 ms",
        "64 bytes from 127.0.0.1: icmp_seq=73 ttl=64 time=0.020 ms",
        "64 bytes from 127.0.0.1: icmp_seq=74 ttl=64 time=0.020 ms",
        ""
      ],
      "moreLine": false
    },
    "state": "running"
  },
  "urn": "linux:jobs:ping:f072dba4-5c93-4136-876d-cf6b8f6cab9e"
}

# kill job (kill -SIGINT)
curl -X POST "${SERVER}/api/v1/jobs/${JOBURN}/kill"

# kill job (kill -SIGKILL)
curl -X POST "${SERVER}/api/v1/jobs/${JOBURN}/kill?force=true"

# delete job
curl -X DELETE "${SERVER}/api/v1/jobs/${JOBURN}"

```

### Operation scp

Copy files from/to other hosts via scp

Dependencies:
- expect
- crypto-tool

```sh
# submit job
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
        "plugin": "linux",
        "operation": "scp",
        "args": ["-h","8.46.188.21","-u","root","-f","/root/k9s_Linux_x86_64.tar.gz","-t","/root/","-m","PUT"],
        "env": ["SCP_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAlX6tkg+bmJ1yE7l1iYJz+FiZN8XWeIPVqbm08dAAAAAAEAAAEAAAAAAAAAGiAYvj77qxKCKsBaxBXaCuIwgqFp7sluRvaj"]
      }]'

[
  {
    "id": "b7422b38-2c50-41c0-b1f1-28a7d1008d39",
    "spec": {
      "args": [
        "-h",
        "8.46.188.21",
        "-u",
        "root",
        "-f",
        "/root/k9s_Linux_x86_64.tar.gz",
        "-t",
        "/root/",
        "-m",
        "PUT"
      ],
      "env": [
        "SCP_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAlX6tkg+bmJ1yE7l1iYJz+FiZN8XWeIPVqbm08dAAAAAAEAAAEAAAAAAAAAGiAYvj77qxKCKsBaxBXaCuIwgqFp7sluRvaj"
      ],
      "operation": "scp",
      "plugin": "linux"
    },
    "status": {
      "state": "waiting"
    },
    "urn": "linux:jobs:scp:b7422b38-2c50-41c0-b1f1-28a7d1008d39"
  }
]

JOBURN="linux:jobs:scp:b7422b38-2c50-41c0-b1f1-28a7d1008d39"

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}"

{
  "id": "b7422b38-2c50-41c0-b1f1-28a7d1008d39",
  "spec": {
    "args": [
      "-h",
      "8.46.188.21",
      "-u",
      "root",
      "-f",
      "/root/k9s_Linux_x86_64.tar.gz",
      "-t",
      "/root/",
      "-m",
      "PUT"
    ],
    "env": [
      "SCP_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAlX6tkg+bmJ1yE7l1iYJz+FiZN8XWeIPVqbm08dAAAAAAEAAAEAAAAAAAAAGiAYvj77qxKCKsBaxBXaCuIwgqFp7sluRvaj"
    ],
    "operation": "scp",
    "plugin": "linux"
  },
  "status": {
    "exitCode": 0,
    "output": {
      "lastLine": 6,
      "lines": [
        "spawn /usr/bin/scp -o StrictHostKeyChecking=no -P 22 /root/k9s_Linux_x86_64.tar.gz root@8.46.188.21:/root/\r",
        "\r",
        "Authorized users only. All activities may be monitored and reported.\r",
        "\rroot@8.46.188.21's password: \r",
        "\rk9s_Linux_x86_64.tar.gz                       100%   15MB 104.3MB/s   00:00    \r",
        ""
      ],
      "moreLine": false
    },
    "state": "exited"
  },
  "urn": "linux:jobs:scp:b7422b38-2c50-41c0-b1f1-28a7d1008d39"
}

# delete job
curl -X DELETE "${SERVER}/api/v1/jobs/${JOBURN}"

```

### Operation write

Write input data to a specific file.

```sh
# submit job
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{"plugin":"linux","operation":"write","args":["-f","/root/write.out","-m","APPEND"]}]'

[
  {
    "id": "2a7dc7ad-ad9d-4190-9d91-09c8cdd1ba75",
    "spec": {
      "args": [
        "-f",
        "/root/write.out",
        "-m",
        "APPEND"
      ],
      "env": null,
      "operation": "write",
      "plugin": "linux"
    },
    "status": {
      "state": "waiting"
    },
    "urn": "linux:jobs:write:2a7dc7ad-ad9d-4190-9d91-09c8cdd1ba75"
  }
]

JOBURN="linux:jobs:write:2a7dc7ad-ad9d-4190-9d91-09c8cdd1ba75"

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}"

{
  "id": "2a7dc7ad-ad9d-4190-9d91-09c8cdd1ba75",
  "spec": {
    "args": [
      "-f",
      "/root/write.out",
      "-m",
      "APPEND"
    ],
    "env": null,
    "operation": "write",
    "plugin": "linux"
  },
  "status": {
    "output": {
      "lastLine": 2,
      "lines": [
        "",
        "> "
      ],
      "moreLine": false
    },
    "state": "running"
  },
  "urn": "linux:jobs:write:2a7dc7ad-ad9d-4190-9d91-09c8cdd1ba75"
}

# input data
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs/${JOBURN}/input" \
     -d '{"expect":"> ","timeout":1,"data":"line1\n"}'

curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs/${JOBURN}/input" \
     -d '{"expect":"> ","timeout":1,"data":"line2\n"}'

curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs/${JOBURN}/input" \
     -d '{"expect":"> ","timeout":1,"data":"EOF\n"}'

curl "${SERVER}/api/v1/jobs/${JOBURN}"

{
  "id": "2a7dc7ad-ad9d-4190-9d91-09c8cdd1ba75",
  "spec": {
    "args": [
      "-f",
      "/root/write.out",
      "-m",
      "APPEND"
    ],
    "env": null,
    "operation": "write",
    "plugin": "linux"
  },
  "status": {
    "exitCode": 0,
    "output": {
      "lastLine": 4,
      "lines": [
        "",
        "> ",
        "> ",
        "> "
      ],
      "moreLine": false
    },
    "state": "exited"
  },
  "urn": "linux:jobs:write:2a7dc7ad-ad9d-4190-9d91-09c8cdd1ba75"
}

```

### Operation cert_check

Check whether certification expire
- exit code 0: not expired
- exit code 1: expired

```sh
# submit job
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{"plugin":"linux","operation":"cert_check","args":["-f","/opt/flexagent/certs/ca.crt","-e","90"]}]'

[
  {
    "id": "42afa8b0-bbbf-4f37-a704-c479371e7a20",
    "spec": {
      "args": [
        "-f",
        "/etc/flexagent/certs/ca.crt",
        "-d",
        "90"
      ],
      "env": null,
      "operation": "cert_check",
      "plugin": "linux"
    },
    "status": {
      "state": "waiting"
    },
    "urn": "linux:jobs:cert_check:42afa8b0-bbbf-4f37-a704-c479371e7a20"
  }
]

JOBURN="linux:jobs:cert_check:42afa8b0-bbbf-4f37-a704-c479371e7a20"

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}"

{
  "id": "42afa8b0-bbbf-4f37-a704-c479371e7a20",
  "spec": {
    "args": [
      "-f",
      "/etc/flexagent/certs/ca.crt",
      "-d",
      "90"
    ],
    "env": null,
    "operation": "cert_check",
    "plugin": "linux"
  },
  "status": {
    "exitCode": 0,
    "output": {
      "lastLine": 2,
      "lines": [
        "Certificate will not expire",
        ""
      ],
      "moreLine": false
    },
    "state": "exited"
  },
  "urn": "linux:jobs:cert_check:42afa8b0-bbbf-4f37-a704-c479371e7a20"
}

# create cronjob
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/cronjobs" \
     -d '[{
       "name": "cert_check_1h",
       "schedule": "0 0 */1 * * *",
       "jobspec": {"plugin":"linux","operation":"cert_check","args":["-f","/etc/flexagent/certs/ca.crt","-d","90"]}
      }]'

CRONID="8df42ceb-34ae-4266-b86b-dcc8fecf1348"

# query cronjob
curl "${SERVER}/api/v1/cronjobs/${CRONID}"
```

### Operation cert_gen

Generte certs

Dependencies:
- openssl
- crypto-tool

```sh
# submit job
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
        "plugin": "linux",
        "operation": "cert_gen",
        "args": ["-n","example.com","-c","/root/certs/client1.crt","-k","/root/certs/client1.key"],
        "env": [
          "CERT_ENV_CA_CERT=LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNzekNDQVpzQ0ZFeEhZWlUwZUNXSGc2OFZqTGZUQy9mOG43dVlNQTBHQ1NxR1NJYjNEUUVCQ3dVQU1CWXgKRkRBU0JnTlZCQU1NQzJWNFlXMXdiR1V1WTI5dE1CNFhEVEl5TURjeU16RXlOVGMwT0ZvWERUTXlNRGN5TURFeQpOVGMwT0Zvd0ZqRVVNQklHQTFVRUF3d0xaWGhoYlhCc1pTNWpiMjB3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBCkE0SUJEd0F3Z2dFS0FvSUJBUUN1ZWthKzJQejJBZE03Y3dOMmwrN3lZbjB6SmVSdVEvZHVsUjhkTHdYN1ZKYTIKTFg4WE9qeDNqNk9kSnVIM00rUlVDOUhNYkk1cXc1dnhScXhUaXRQc2UxZlB4U2J4dm40dFpYQWZjMFZ4bjlmdApWdjhNTTNqajMxSS9nWXdjM3lYRnkxNm9tY3BueFhqTXNoM0tJZDYxMUJMMXBNbGtIRUtxMzduSTdmSmR1YTdhClJWajdUV1BoOEJWd0w4N201ckg1YkhzdW43a0JKTnZKbFVCQ3hiRnViTnRlUzVVbnhHQWJvdnFzZkl4TjJJdVgKZUVKUE9NQ3Avc2djbWZNd3NIODl0VVJrYjNQVlZFUmYrT2t6VC9OSUVVMTZsakVWSTZGbFFvMlN4NmZZNXUwcwplOFA3Nm9UZFZaK1Bmc1NmQTJOMzc5Qlo5RUhvK2t2TkJTQ1RsMEpwQWdNQkFBRXdEUVlKS29aSWh2Y05BUUVMCkJRQURnZ0VCQURaWExVdlRLY2NzRWJwbnNMRTgwUkNOZm8zMWE0dzB3YnJ3V0hxVy9pZjlDRXY1ckw4ampxRHIKUU9JcnRYSUtGSndRVTJiOVhsQS9IU1JwNE9DbEhobWhUMU80N2NCK2d3U25KMFdtNHoxd3hNR0VyWkZ4REFEMwpTL1hoSnRvT1p6bDNPb1ZnVWphL1ZacXZ4dWxzeWZnSi83YlhMU1l5WnlrWFFnNmlRR2doRGFJcFlnbWgzK2JOCmVacitoVUxZWjAwRkdYUkRPYS9RUVNtOWUyb1pLb0tBdXFTRFdQTUtmM2Q1bWw5NlpoVWRyQ2x2dnZlT01WVmMKL01iS3dYUnZSNzQvbXdnVXptbCtWM2xNWkZ0bG1LRG1UYkdPUGZ2RDU2ajcrK2NEcjZSa0laSUdFWmFyMDE5cwo3UkdsVTZZNGk5UW52ZitXYUM0UmxnN1V5cytLMW1VPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
          "CERT_ENV_CA_KEY=LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcm5wR3Z0ajg5Z0hUTzNNRGRwZnU4bUo5TXlYa2JrUDNicFVmSFM4RisxU1d0aTEvCkZ6bzhkNCtqblNiaDl6UGtWQXZSekd5T2FzT2I4VWFzVTRyVDdIdFh6OFVtOGI1K0xXVndIM05GY1ovWDdWYi8KRERONDQ5OVNQNEdNSE44bHhjdGVxSm5LWjhWNHpMSWR5aUhldGRRUzlhVEpaQnhDcXQrNXlPM3lYYm11MmtWWQorMDFqNGZBVmNDL081dWF4K1d4N0xwKzVBU1RieVpWQVFzV3hibXpiWGt1Vko4UmdHNkw2ckh5TVRkaUxsM2hDClR6akFxZjdJSEpuek1MQi9QYlZFWkc5ejFWUkVYL2pwTTAvelNCRk5lcFl4RlNPaFpVS05rc2VuMk9idExIdkQKKytxRTNWV2ZqMzdFbndOamQrL1FXZlJCNlBwTHpRVWdrNWRDYVFJREFRQUJBb0lCQURURml3eENPUkIzTU5wMApTUUhmcGtmdGZneXdVMHoraGFJT3ZHQnBUcGZiMTlHMkpSQnpic0tDMFd2QXpPdWw2Rk43VjdOS3lqQ0VoWEFPCmZpU3VncmF0LzdzNTJET1orRFBtMEdqU3hwZzlUbVBjd2p1QWZmbitHN1JWakhsWERPZXNRUzhoaE5TcUo0VlMKQWl5VzNmT0U5MDhRemxEWE0xclhYQ3ZudkR5WVlyM1JQV3FKd0JCL3h4bGF6SXBtYjdxRzJwOGJua3NHVlR1dApQcngzb0hKNllDMUlILyt0bGFsSjkyNHp0WmNGUWVYWU1rZUxkVmtPYnQ3R2dLc3E4SmZud25iNVF2Q0ZsTmMvClVXZWxtZi9kRUpqN0ZaQktMMWtDcld6SFJwQXZ5em9Sc0hMK0NkSVljNytwMVJxd0Q2emVYRGNLWFd3aHJJWTcKamgyazNlRUNnWUVBMzFvR0hwbVBCV0E1cjBYZ1c4Z21aTmNXMVpTWGFmeVRHSG9HREN3elV3TG5UNHZpcU5nQwpZWDltSFppL25uYnVPN0JHWW4wS1pNcmE5WjJ1Rm5SM2toSHQ2anJld3p4ejhjaWw5WkY5TWlSQXRDWFZ4bjc4CkFBTUxFemRteVFNSW9Bd1Bnb2FFcElJY1krLzJOVk9Ia0R1bWNTUXNZazJEWWpCT2hNd1hMOTBDZ1lFQXgvdGEKT3UxV3p6TUN5SUxIYmt6ZTZIWnQzRWkzaFVDWEIyWHluZ3ZMeHQvdEpESmp5dEFQYWs5RS9aN2VidUJPY1Zhcwo1b09wWUFZSzVia3puL1kzTEZCdFFMWGdiWVVPMmJqdE9hVDh3aWQ0REZNN1ZYQ2draE94aTF3cmlrMFlNWkZrCkZqWG0zRDltdUtvR21ZUkxYUlBjeE9KZVdqY2doMmg3K0YwcitmMENnWUEzVjhneVp0eGdlYUp4Z3NBQUhnMGQKYVlwMzY3VEZCMWV2ZGZUdnFUZ2lkcEs0VERJaW9qdWN5d09UaTlqWFBDTDEyVXpuZEpKUnZVNGFGRE1oejBRZApocUhNSzBBdFlscGNhOXByaWR4YXcwN2hGSXJ1LzJJVDRxMG8yczUyT25FMXJ5ZGNzVlpHcVJLOTFLVE9POTlZClp0OXNJNGwzNWpzSzVtdGVUbS9rWlFLQmdRQ0VQd0QxVlB5Q090Nk5VSWFudDJmMVhGUGNSNjR0RFlDU29PVm4KaEs5MlRhRFp2Z1RtR3Q4RzAzTHhNVDB4SDE1Z2J3d1p5Rm1hcVlSTlZFTUNkbVVZQmZ1cHZseXlzRG9ZMnNUdAp5T0JwV0laM3lCYkZzcHhNM1g4Y2hKQTZmaThRb0hBS2pBeWwrN3RuUlBEbVZta3NIVFZ5Y2F3cGhxa1pRb3d3CnV4U1kxUUtCZ1FDeE1GQTQrT3VYSDFOVXZJcUJVeUtpVy9rV3JEd25aOGdoUDVBSlBBbXBydDBNdHlydXlBRS8KSmJzRWl0clRZU1Q5N1FTREhCU1M5aFM1SlQ0bVhUbEFyRllSdWRvME5tZjNBVUlHSWJDRXlsbDJtRHdXZ2xxZApDRUFXREd5ZlBiSWNZUkQzc015T2wvSk9FNHFGMUdreXJoUEp0UnpBTFRja0dpTm0wQ2VNNHc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=",
          "CERT_ENV_CA_KEY_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAlX6tkg+bmJ1zo4zBhwnh5JOqVXRiCE1oVUjOHYAAAAAAEAAAEAAAAAAAAAE+c9UPnec1kgG/LxopgtO7x0qWM="
        ]
      }]'

JOBURN="linux:jobs:cert_gen:60ca82fe-b2b2-4a34-b9e5-6ee13c181c90"

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}"


{
  "id": "60ca82fe-b2b2-4a34-b9e5-6ee13c181c90",
  "spec": {
    "args": [
      "-n",
      "example.com",
      "-c",
      "/root/certs/client1.crt",
      "-k",
      "/root/certs/client1.key"
    ],
    "env": [
      "CERT_ENV_CA_CERT=LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNzekNDQVpzQ0ZFeEhZWlUwZUNXSGc2OFZqTGZUQy9mOG43dVlNQTBHQ1NxR1NJYjNEUUVCQ3dVQU1CWXgKRkRBU0JnTlZCQU1NQzJWNFlXMXdiR1V1WTI5dE1CNFhEVEl5TURjeU16RXlOVGMwT0ZvWERUTXlNRGN5TURFeQpOVGMwT0Zvd0ZqRVVNQklHQTFVRUF3d0xaWGhoYlhCc1pTNWpiMjB3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBCkE0SUJEd0F3Z2dFS0FvSUJBUUN1ZWthKzJQejJBZE03Y3dOMmwrN3lZbjB6SmVSdVEvZHVsUjhkTHdYN1ZKYTIKTFg4WE9qeDNqNk9kSnVIM00rUlVDOUhNYkk1cXc1dnhScXhUaXRQc2UxZlB4U2J4dm40dFpYQWZjMFZ4bjlmdApWdjhNTTNqajMxSS9nWXdjM3lYRnkxNm9tY3BueFhqTXNoM0tJZDYxMUJMMXBNbGtIRUtxMzduSTdmSmR1YTdhClJWajdUV1BoOEJWd0w4N201ckg1YkhzdW43a0JKTnZKbFVCQ3hiRnViTnRlUzVVbnhHQWJvdnFzZkl4TjJJdVgKZUVKUE9NQ3Avc2djbWZNd3NIODl0VVJrYjNQVlZFUmYrT2t6VC9OSUVVMTZsakVWSTZGbFFvMlN4NmZZNXUwcwplOFA3Nm9UZFZaK1Bmc1NmQTJOMzc5Qlo5RUhvK2t2TkJTQ1RsMEpwQWdNQkFBRXdEUVlKS29aSWh2Y05BUUVMCkJRQURnZ0VCQURaWExVdlRLY2NzRWJwbnNMRTgwUkNOZm8zMWE0dzB3YnJ3V0hxVy9pZjlDRXY1ckw4ampxRHIKUU9JcnRYSUtGSndRVTJiOVhsQS9IU1JwNE9DbEhobWhUMU80N2NCK2d3U25KMFdtNHoxd3hNR0VyWkZ4REFEMwpTL1hoSnRvT1p6bDNPb1ZnVWphL1ZacXZ4dWxzeWZnSi83YlhMU1l5WnlrWFFnNmlRR2doRGFJcFlnbWgzK2JOCmVacitoVUxZWjAwRkdYUkRPYS9RUVNtOWUyb1pLb0tBdXFTRFdQTUtmM2Q1bWw5NlpoVWRyQ2x2dnZlT01WVmMKL01iS3dYUnZSNzQvbXdnVXptbCtWM2xNWkZ0bG1LRG1UYkdPUGZ2RDU2ajcrK2NEcjZSa0laSUdFWmFyMDE5cwo3UkdsVTZZNGk5UW52ZitXYUM0UmxnN1V5cytLMW1VPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
      "CERT_ENV_CA_KEY=LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcm5wR3Z0ajg5Z0hUTzNNRGRwZnU4bUo5TXlYa2JrUDNicFVmSFM4RisxU1d0aTEvCkZ6bzhkNCtqblNiaDl6UGtWQXZSekd5T2FzT2I4VWFzVTRyVDdIdFh6OFVtOGI1K0xXVndIM05GY1ovWDdWYi8KRERONDQ5OVNQNEdNSE44bHhjdGVxSm5LWjhWNHpMSWR5aUhldGRRUzlhVEpaQnhDcXQrNXlPM3lYYm11MmtWWQorMDFqNGZBVmNDL081dWF4K1d4N0xwKzVBU1RieVpWQVFzV3hibXpiWGt1Vko4UmdHNkw2ckh5TVRkaUxsM2hDClR6akFxZjdJSEpuek1MQi9QYlZFWkc5ejFWUkVYL2pwTTAvelNCRk5lcFl4RlNPaFpVS05rc2VuMk9idExIdkQKKytxRTNWV2ZqMzdFbndOamQrL1FXZlJCNlBwTHpRVWdrNWRDYVFJREFRQUJBb0lCQURURml3eENPUkIzTU5wMApTUUhmcGtmdGZneXdVMHoraGFJT3ZHQnBUcGZiMTlHMkpSQnpic0tDMFd2QXpPdWw2Rk43VjdOS3lqQ0VoWEFPCmZpU3VncmF0LzdzNTJET1orRFBtMEdqU3hwZzlUbVBjd2p1QWZmbitHN1JWakhsWERPZXNRUzhoaE5TcUo0VlMKQWl5VzNmT0U5MDhRemxEWE0xclhYQ3ZudkR5WVlyM1JQV3FKd0JCL3h4bGF6SXBtYjdxRzJwOGJua3NHVlR1dApQcngzb0hKNllDMUlILyt0bGFsSjkyNHp0WmNGUWVYWU1rZUxkVmtPYnQ3R2dLc3E4SmZud25iNVF2Q0ZsTmMvClVXZWxtZi9kRUpqN0ZaQktMMWtDcld6SFJwQXZ5em9Sc0hMK0NkSVljNytwMVJxd0Q2emVYRGNLWFd3aHJJWTcKamgyazNlRUNnWUVBMzFvR0hwbVBCV0E1cjBYZ1c4Z21aTmNXMVpTWGFmeVRHSG9HREN3elV3TG5UNHZpcU5nQwpZWDltSFppL25uYnVPN0JHWW4wS1pNcmE5WjJ1Rm5SM2toSHQ2anJld3p4ejhjaWw5WkY5TWlSQXRDWFZ4bjc4CkFBTUxFemRteVFNSW9Bd1Bnb2FFcElJY1krLzJOVk9Ia0R1bWNTUXNZazJEWWpCT2hNd1hMOTBDZ1lFQXgvdGEKT3UxV3p6TUN5SUxIYmt6ZTZIWnQzRWkzaFVDWEIyWHluZ3ZMeHQvdEpESmp5dEFQYWs5RS9aN2VidUJPY1Zhcwo1b09wWUFZSzVia3puL1kzTEZCdFFMWGdiWVVPMmJqdE9hVDh3aWQ0REZNN1ZYQ2draE94aTF3cmlrMFlNWkZrCkZqWG0zRDltdUtvR21ZUkxYUlBjeE9KZVdqY2doMmg3K0YwcitmMENnWUEzVjhneVp0eGdlYUp4Z3NBQUhnMGQKYVlwMzY3VEZCMWV2ZGZUdnFUZ2lkcEs0VERJaW9qdWN5d09UaTlqWFBDTDEyVXpuZEpKUnZVNGFGRE1oejBRZApocUhNSzBBdFlscGNhOXByaWR4YXcwN2hGSXJ1LzJJVDRxMG8yczUyT25FMXJ5ZGNzVlpHcVJLOTFLVE9POTlZClp0OXNJNGwzNWpzSzVtdGVUbS9rWlFLQmdRQ0VQd0QxVlB5Q090Nk5VSWFudDJmMVhGUGNSNjR0RFlDU29PVm4KaEs5MlRhRFp2Z1RtR3Q4RzAzTHhNVDB4SDE1Z2J3d1p5Rm1hcVlSTlZFTUNkbVVZQmZ1cHZseXlzRG9ZMnNUdAp5T0JwV0laM3lCYkZzcHhNM1g4Y2hKQTZmaThRb0hBS2pBeWwrN3RuUlBEbVZta3NIVFZ5Y2F3cGhxa1pRb3d3CnV4U1kxUUtCZ1FDeE1GQTQrT3VYSDFOVXZJcUJVeUtpVy9rV3JEd25aOGdoUDVBSlBBbXBydDBNdHlydXlBRS8KSmJzRWl0clRZU1Q5N1FTREhCU1M5aFM1SlQ0bVhUbEFyRllSdWRvME5tZjNBVUlHSWJDRXlsbDJtRHdXZ2xxZApDRUFXREd5ZlBiSWNZUkQzc015T2wvSk9FNHFGMUdreXJoUEp0UnpBTFRja0dpTm0wQ2VNNHc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=",
      "CERT_ENV_CA_KEY_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAlX6tkg+bmJ1zo4zBhwnh5JOqVXRiCE1oVUjOHYAAAAAAEAAAEAAAAAAAAAE+c9UPnec1kgG/LxopgtO7x0qWM="
    ],
    "operation": "cert_gen",
    "plugin": "linux"
  },
  "status": {
    "exitCode": 0,
    "output": {
      "lastLine": 1,
      "lines": [
        ""
      ],
      "moreLine": false
    },
    "state": "exited"
  },
  "urn": "linux:jobs:cert_gen:60ca82fe-b2b2-4a34-b9e5-6ee13c181c90"
}
```
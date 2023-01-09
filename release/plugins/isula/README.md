# isula plugin

## Probes

### Probe service_status

```sh
/opt/flexagent/plugins/isula/probes/service_status.sh 

isulad: [active]
```

## Operations

Test on local:

```sh
# test on local
SERVER="http://127.0.0.1:18080"
alias JSON="python -c \"import json; import sys; print(json.dumps(json.loads(sys.stdin.read()), indent=4, ensure_ascii=False));\""
```

### Operation registry_cfg

Configure image registry

```sh
# submit jobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
          "plugin": "isula",
          "operation": "registry_cfg",
          "args": ["-a","192.168.1.51","-n","harbor.dept.example.com","-u","admin"],
          "env": [
              "IMAGE_REPO_CERT=LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZ1ekNDQTZPZ0F3SUJBZ0lKQUxNTlFWQWladTk0TUEwR0NTcUdTSWIzRFFFQkN3VUFNSFF4Q3pBSkJnTlYKQkFZVEFrTk9NUkF3RGdZRFZRUUlEQWRUYVdOb2RXRnVNUkF3RGdZRFZRUUhEQWREYUdWdVoyUjFNUTh3RFFZRApWUVFLREFaSWRXRjNaV2t4RURBT0JnTlZCQXNNQjFOMGIzSmhaMlV4SGpBY0JnTlZCQU1NRldoaGNtSnZjaTVrCmJXVXVhSFZoZDJWcExtTnZiVEFlRncweU1UQTNNamd4TXpJMU5EZGFGdzB6TVRBM01qWXhNekkxTkRkYU1IUXgKQ3pBSkJnTlZCQVlUQWtOT01SQXdEZ1lEVlFRSURBZFRhV05vZFdGdU1SQXdEZ1lEVlFRSERBZERhR1Z1WjJSMQpNUTh3RFFZRFZRUUtEQVpJZFdGM1pXa3hFREFPQmdOVkJBc01CMU4wYjNKaFoyVXhIakFjQmdOVkJBTU1GV2hoCmNtSnZjaTVrYldVdWFIVmhkMlZwTG1OdmJUQ0NBaUl3RFFZSktvWklodmNOQVFFQkJRQURnZ0lQQURDQ0Fnb0MKZ2dJQkFLeWVzK2psckRLdlREUWE3Q1QwRmhBYmZrSURJT1Q4ZE95OTBaL0VkU09GWUx3c1Q2WGRtNnZIUWk5dQo4NmYrd1RZUzhBZHBFNU05a0p4S0JTYzFrS2lqbGFEYlBBdjBabkg0dGk3Z1pjZWhqd2tJUkFZdTc0RzZqM2dHClFETlRDOFZWRllETVdsdnlJc3FDbjZjUWFMNGFwOWgycGM0YTk1WDN2K2xqeTdKV2F1T0p1V0pYTUdsNFRVN3UKQUJkVEErMjAzdCtucWsvbzJxVW5mbUl3VGxHQ0EreXJlckVOQTdwa283ZTRtMjJqY245bmFyUW9QcUpwSTY5bgpRWkhtQ3VLL2xCSE5iVkVkRlU0WVFtbkpkQ05nNFR5cWhtbFpqTTViY2tUbWdVK3JDM2ZtUFZFVHU0bm41KzJhClhTcHc2QlB4c0xTTkJURVpVUWtLM1F1OW1JYjNxdHc1YkxWY2h5OHBpMnhqZE1xZVJUOU1MMlI1RzJyM2xIaDkKRFBuM1NFa1gvalFNUE9QMkNwekhqcVlaZ3RmWnNFYnJUdk4rYWh5bWpoTnNBUFFKYXR3WVZJMldXUGdlU2VBZwpYN3VPVzRwS1FtU04vWnJCV3RGVkI5MTV2NjBJbWs0cmtVR3g5Tnl2OUJWc1EwVm9XNVFUUHZXZ0l0WmJvYVZ4Cktnc0RCQkl1V3BGYW9lOVFiMlMvREJ4MktYQTN4dDVvekoxeTVVMVd5VFllaFhhdzQzcW9EaFNKa05ObUlsSEoKTjc3cngrMERlcHI3V2tERzRBV3hKSjhITU4xZGNiMFlkR3NHdTFaL1QwZlQvTXJkcDJwcW9Oem9yYTlqVW1iTQpFQXI2Vk00d25vSi83dHRpWnV6ZFFmWUlHTEZQZk5BZFZkYXA2c2FhNFpRZUJHYVJBZ01CQUFHalVEQk9NQjBHCkExVWREZ1FXQkJUMnFFVVkwOHZOSUg3dFpkdk5NU0RZM0ovV2VEQWZCZ05WSFNNRUdEQVdnQlQycUVVWTA4dk4KSUg3dFpkdk5NU0RZM0ovV2VEQU1CZ05WSFJNRUJUQURBUUgvTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElDQVFCZgpvYm8wS3U5NVNYQmZoeXBmcHFSSEdlNkpFU0FZM0FhV2pWeW9VRkV2SmFUNTdTY1QzNi9PNldnNFdiRnJGY291Ck9TRHBUVlVXNjZ3dUJTdVg2S1N0QVJCMkJZcm1MMXY1MFVHbVhRUTZmR2g1QTYxYTVCblVCdXhQTjJSUEhmZ2MKenBvTHFmZkZBcmphUWZyTEpXa1gzWGxGbXJrbjVuSkl2WWF3WmNvRHRiZ3ArOHlEaDVnc2VnRjQyaUdnczFDQgpacnZVVXBCZ0JueWVrWkNFbGcwVjRQVnhiMXVhVGxYZVpIR1hOU1l3S3Y1Q1MvcTEwbnVkY2d1Rm5kNW1KdlFwCk1PSHFGMGZrcGdZV0FWRGJXSEFyUDdRYUI5d0RDUTVxYjJlY24rY0dNejZEbDlCdVZCYnMzZHJkdU5KSU5CL1YKYTRid3FDLzFhemdRNUczRGJyTDVVRUgzclhxVzZIQkJLV0J3VFp3Q1BKSkI0ekRlVXQwRi9NcUNxbDlYdGszeApocUJDZXcxSmtra1FOMmtsYVFsK2twb3NpaU5RNUlYZFJmVzF1Q3ZYVit5TllZRVlCY0VDWjNYZ1lvVlNsM0lsCmJsWnNxSDlYaHp0dmJNZHp1VlhFOHZyYll0QWRybHA2ekdGeTI5NTg1Yld2SG42MytJUEhCN284STZCKzEwL2EKWDFiWTV2bjZNVkVtTU5zbDh4MFdJOEo1ekI3endWYkMxZlV5SUY4NnYzZmRkMHh4VFpucE4xdVEzWjV1QTRWVQpPaVJjRTB4VGl0RGd4MUVFa1BoTmk3cjl0SkIzNW14VFBHeEl0ejE2TVk3VGxqS0NNQTVOU0xzWGZOSnpvTTlkClBpUEk5Q3R6K3hOWTZzcWl3b3FEeUFxd0JCclZBQkhWWnpHZ3UvcWd6UT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
              "IMAGE_REPO_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAl5XukHCU1z+AwKQrUcLZ6r54g+cKxXooZjLddAAAAAAAEAAAEAAAAAAAAAGpr1Ln8ar6Onj/5E2upEJrKMJVDnDGtXFAtK"
          ]
         }]'

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}" | JSON

{
    "id": "c2483ad2-6498-4979-9cc1-ac42a63cd0fd",
    "spec": {
        "args": [
            "-a",
            "192.168.1.51",
            "-n",
            "harbor.dept.example.com",
            "-u",
            "admin"
        ],
        "env": [
            "IMAGE_REPO_CERT=LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZ1ekNDQTZPZ0F3SUJBZ0lKQUxNTlFWQWladTk0TUEwR0NTcUdTSWIzRFFFQkN3VUFNSFF4Q3pBSkJnTlYKQkFZVEFrTk9NUkF3RGdZRFZRUUlEQWRUYVdOb2RXRnVNUkF3RGdZRFZRUUhEQWREYUdWdVoyUjFNUTh3RFFZRApWUVFLREFaSWRXRjNaV2t4RURBT0JnTlZCQXNNQjFOMGIzSmhaMlV4SGpBY0JnTlZCQU1NRldoaGNtSnZjaTVrCmJXVXVhSFZoZDJWcExtTnZiVEFlRncweU1UQTNNamd4TXpJMU5EZGFGdzB6TVRBM01qWXhNekkxTkRkYU1IUXgKQ3pBSkJnTlZCQVlUQWtOT01SQXdEZ1lEVlFRSURBZFRhV05vZFdGdU1SQXdEZ1lEVlFRSERBZERhR1Z1WjJSMQpNUTh3RFFZRFZRUUtEQVpJZFdGM1pXa3hFREFPQmdOVkJBc01CMU4wYjNKaFoyVXhIakFjQmdOVkJBTU1GV2hoCmNtSnZjaTVrYldVdWFIVmhkMlZwTG1OdmJUQ0NBaUl3RFFZSktvWklodmNOQVFFQkJRQURnZ0lQQURDQ0Fnb0MKZ2dJQkFLeWVzK2psckRLdlREUWE3Q1QwRmhBYmZrSURJT1Q4ZE95OTBaL0VkU09GWUx3c1Q2WGRtNnZIUWk5dQo4NmYrd1RZUzhBZHBFNU05a0p4S0JTYzFrS2lqbGFEYlBBdjBabkg0dGk3Z1pjZWhqd2tJUkFZdTc0RzZqM2dHClFETlRDOFZWRllETVdsdnlJc3FDbjZjUWFMNGFwOWgycGM0YTk1WDN2K2xqeTdKV2F1T0p1V0pYTUdsNFRVN3UKQUJkVEErMjAzdCtucWsvbzJxVW5mbUl3VGxHQ0EreXJlckVOQTdwa283ZTRtMjJqY245bmFyUW9QcUpwSTY5bgpRWkhtQ3VLL2xCSE5iVkVkRlU0WVFtbkpkQ05nNFR5cWhtbFpqTTViY2tUbWdVK3JDM2ZtUFZFVHU0bm41KzJhClhTcHc2QlB4c0xTTkJURVpVUWtLM1F1OW1JYjNxdHc1YkxWY2h5OHBpMnhqZE1xZVJUOU1MMlI1RzJyM2xIaDkKRFBuM1NFa1gvalFNUE9QMkNwekhqcVlaZ3RmWnNFYnJUdk4rYWh5bWpoTnNBUFFKYXR3WVZJMldXUGdlU2VBZwpYN3VPVzRwS1FtU04vWnJCV3RGVkI5MTV2NjBJbWs0cmtVR3g5Tnl2OUJWc1EwVm9XNVFUUHZXZ0l0WmJvYVZ4Cktnc0RCQkl1V3BGYW9lOVFiMlMvREJ4MktYQTN4dDVvekoxeTVVMVd5VFllaFhhdzQzcW9EaFNKa05ObUlsSEoKTjc3cngrMERlcHI3V2tERzRBV3hKSjhITU4xZGNiMFlkR3NHdTFaL1QwZlQvTXJkcDJwcW9Oem9yYTlqVW1iTQpFQXI2Vk00d25vSi83dHRpWnV6ZFFmWUlHTEZQZk5BZFZkYXA2c2FhNFpRZUJHYVJBZ01CQUFHalVEQk9NQjBHCkExVWREZ1FXQkJUMnFFVVkwOHZOSUg3dFpkdk5NU0RZM0ovV2VEQWZCZ05WSFNNRUdEQVdnQlQycUVVWTA4dk4KSUg3dFpkdk5NU0RZM0ovV2VEQU1CZ05WSFJNRUJUQURBUUgvTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElDQVFCZgpvYm8wS3U5NVNYQmZoeXBmcHFSSEdlNkpFU0FZM0FhV2pWeW9VRkV2SmFUNTdTY1QzNi9PNldnNFdiRnJGY291Ck9TRHBUVlVXNjZ3dUJTdVg2S1N0QVJCMkJZcm1MMXY1MFVHbVhRUTZmR2g1QTYxYTVCblVCdXhQTjJSUEhmZ2MKenBvTHFmZkZBcmphUWZyTEpXa1gzWGxGbXJrbjVuSkl2WWF3WmNvRHRiZ3ArOHlEaDVnc2VnRjQyaUdnczFDQgpacnZVVXBCZ0JueWVrWkNFbGcwVjRQVnhiMXVhVGxYZVpIR1hOU1l3S3Y1Q1MvcTEwbnVkY2d1Rm5kNW1KdlFwCk1PSHFGMGZrcGdZV0FWRGJXSEFyUDdRYUI5d0RDUTVxYjJlY24rY0dNejZEbDlCdVZCYnMzZHJkdU5KSU5CL1YKYTRid3FDLzFhemdRNUczRGJyTDVVRUgzclhxVzZIQkJLV0J3VFp3Q1BKSkI0ekRlVXQwRi9NcUNxbDlYdGszeApocUJDZXcxSmtra1FOMmtsYVFsK2twb3NpaU5RNUlYZFJmVzF1Q3ZYVit5TllZRVlCY0VDWjNYZ1lvVlNsM0lsCmJsWnNxSDlYaHp0dmJNZHp1VlhFOHZyYll0QWRybHA2ekdGeTI5NTg1Yld2SG42MytJUEhCN284STZCKzEwL2EKWDFiWTV2bjZNVkVtTU5zbDh4MFdJOEo1ekI3endWYkMxZlV5SUY4NnYzZmRkMHh4VFpucE4xdVEzWjV1QTRWVQpPaVJjRTB4VGl0RGd4MUVFa1BoTmk3cjl0SkIzNW14VFBHeEl0ejE2TVk3VGxqS0NNQTVOU0xzWGZOSnpvTTlkClBpUEk5Q3R6K3hOWTZzcWl3b3FEeUFxd0JCclZBQkhWWnpHZ3UvcWd6UT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
            "IMAGE_REPO_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAl5XukHCU1z+AwKQrUcLZ6r54g+cKxXooZjLddAAAAAAAEAAAEAAAAAAAAAGpr1Ln8ar6Onj/5E2upEJrKMJVDnDGtXFAtK"
        ],
        "operation": "registry_cfg",
        "plugin": "isula"
    },
    "status": {
        "exitCode": 0,
        "output": {
            "lastLine": 3,
            "lines": [
                "Login Succeeded",
                "Restart isulad ... active",
                ""
            ],
            "moreLine": false
        },
        "state": "exited"
    },
    "urn": "isula:jobs:registry_cfg:c2483ad2-6498-4979-9cc1-ac42a63cd0fd"
}
```

### Operation registry_del

Remove image registry

```sh
# submit jobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
          "plugin": "isula",
          "operation": "registry_del",
          "args": ["-r","harbor.dept.example.com"],
          "env": []
         }]'

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}" | JSON

{
    "id": "840c5842-97f6-4837-a66a-1c6c3ffa58bb",
    "spec": {
        "args": [
            "-r",
            "harbor.dept.example.com"
        ],
        "env": [],
        "operation": "registry_del",
        "plugin": "isula"
    },
    "status": {
        "exitCode": 0,
        "output": {
            "lastLine": 3,
            "lines": [
                "Logout Succeeded",
                "Restart isulad ... active",
                ""
            ],
            "moreLine": false
        },
        "state": "exited"
    },
    "urn": "isula:jobs:registry_del:840c5842-97f6-4837-a66a-1c6c3ffa58bb"
}
```

### Operation image_pull

Pull images

```sh
# submit jobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
          "plugin": "isula",
          "operation": "image_pull",
          "args": ["nginx:1.23","redis:7.0.3"],
          "env": []
         }]'

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}" | JSON

{
    "id": "44969575-2856-4d5c-aa01-501da10b6180",
    "spec": {
        "args": [
            "nginx:1.23",
            "redis:7.0.3"
        ],
        "env": [],
        "operation": "image_pull",
        "plugin": "isula"
    },
    "status": {
        "exitCode": 0,
        "output": {
            "lastLine": 5,
            "lines": [
                "Image \"nginx:1.23\" pulling",
                "Image \"41b0e86104ba681811bf60b4d6970ed24dd59e282b36c352b8a55823bbb5e14a\" pulled",
                "Image \"redis:7.0.3\" pulling",
                "Image \"3534610348b5abc4a6f6f7c314b5884a51c251e50f0038cb74d28c08cc7dd2a0\" pulled",
                ""
            ],
            "moreLine": false
        },
        "state": "exited"
    },
    "urn": "isula:jobs:image_pull:44969575-2856-4d5c-aa01-501da10b6180"
}
```

### Operation image_remove

Remove images

```sh
# submit jobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
          "plugin": "isula",
          "operation": "image_remove",
          "args": ["nginx:1.23","redis:7.0.3"],
          "env": []
         }]'

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}" | JSON


{
    "id": "0f65d587-6daa-466b-8efd-7ec558916c3c",
    "spec": {
        "args": [
            "nginx:1.23",
            "redis:7.0.3"
        ],
        "env": [],
        "operation": "image_remove",
        "plugin": "isula"
    },
    "status": {
        "exitCode": 0,
        "output": {
            "lastLine": 3,
            "lines": [
                "Image \"nginx:1.23\" removed",
                "Image \"redis:7.0.3\" removed",
                ""
            ],
            "moreLine": false
        },
        "state": "exited"
    },
    "urn": "isula:jobs:image_remove:0f65d587-6daa-466b-8efd-7ec558916c3c"
}
```

### Operation image_sync

Search and sync images from specific registry

```sh
# submit jobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
          "plugin": "isula",
          "operation": "image_sync",
          "args": ["-r", "harbor.dept.example.com", "-u", "admin", "-k", "nginx,redis", "-t", "1.23,7.0.3"],
          "env": ["IMAGE_REPO_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAl5XukHCU1z+AwKQrUcLZ6r54g+cKxXooZjLddAAAAAAAEAAAEAAAAAAAAAGpr1Ln8ar6Onj/5E2upEJrKMJVDnDGtXFAtK"]
         }]'

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}" | JSON

{
    "id": "ab1ede23-be57-4e41-a551-a8c6845122b5",
    "spec": {
        "args": [
            "-r",
            "harbor.dept.example.com",
            "-u",
            "admin",
            "-k",
            "nginx,redis",
            "-t",
            "1.23,7.0.3"
        ],
        "env": [
            "IMAGE_REPO_CIPHER_PASS=AAAAAgAAAAAAAAABAAAAAAAAAAl5XukHCU1z+AwKQrUcLZ6r54g+cKxXooZjLddAAAAAAAEAAAEAAAAAAAAAGpr1Ln8ar6Onj/5E2upEJrKMJVDnDGtXFAtK"
        ],
        "operation": "image_sync",
        "plugin": "isula"
    },
    "status": {
        "exitCode": 0,
        "output": {
            "lastLine": 5,
            "lines": [
                "Image \"nginx:1.23\" pulling",
                "Image \"41b0e86104ba681811bf60b4d6970ed24dd59e282b36c352b8a55823bbb5e14a\" pulled",
                "Image \"redis:7.0.3\" pulling",
                "Image \"3534610348b5abc4a6f6f7c314b5884a51c251e50f0038cb74d28c08cc7dd2a0\" pulled",
                ""
            ],
            "moreLine": false
        },
        "state": "exited"
    },
    "urn": "isula:jobs:image_sync:ab1ede23-be57-4e41-a551-a8c6845122b5"
}

```


### Operation image_clean

Clean images not used in last x days

```sh
# submit jobs
curl -H 'Accept:application/json' -H 'Content-Type:application/json;charset=utf8' \
     -X POST "${SERVER}/api/v1/jobs" \
     -d '[{
          "plugin": "isula",
          "operation": "image_clean",
          "args": ["-d", "30", "-e", "nginx"],
          "env": []
         }]'

# query job
curl "${SERVER}/api/v1/jobs/${JOBURN}" | JSON

{
    "id": "329ded3b-3190-43b5-a9d5-42084faab87e",
    "spec": {
        "args": [
            "-d",
            "30",
            "-e",
            "nginx"
        ],
        "env": [],
        "operation": "image_clean",
        "plugin": "isula"
    },
    "status": {
        "exitCode": 0,
        "output": {
            "lastLine": 4,
            "lines": [
                "Image \"3534610348b5\" removed",
                "Image \"8234082ee653\" removed",
                "Image \"6c0150b56307\" removed",
                ""
            ],
            "moreLine": false
        },
        "state": "exited"
    },
    "urn": "isula:jobs:image_clean:329ded3b-3190-43b5-a9d5-42084faab87e"
}
```
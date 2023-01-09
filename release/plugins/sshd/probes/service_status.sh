#!/bin/bash

SERVICE="sshd"

STATUS=$(systemctl is-active ${SERVICE})
printf "${SERVICE}: [${STATUS}]"

if [ "${STATUS}" != "active" ]; then
  exit ${STATUS_CRITICAL}
fi

exit ${STATUS_HEALTHY}

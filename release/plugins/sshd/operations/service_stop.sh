#!/bin/bash

SERVICE="sshd"

STATUS=$(systemctl is-active ${SERVICE})
if [ "${STATUS}" != "inactive" ]; then
  systemctl stop ${SERVICE}
  sleep 1
fi

systemctl status ${SERVICE}

STATUS=$(systemctl is-active ${SERVICE})
if [ "${STATUS}" != "inactive" ]; then
    exit -1
fi

exit 0
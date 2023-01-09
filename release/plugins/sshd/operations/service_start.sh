#!/bin/bash

SERVICE="sshd"

STATUS=$(systemctl is-active ${SERVICE})
if [ "${STATUS}" != "active" ]; then
  systemctl start ${SERVICE}
  sleep 1
fi

systemctl status ${SERVICE}

STATUS=$(systemctl is-active ${SERVICE})
if [ "${STATUS}" != "active" ]; then
    exit -1
fi

exit 0
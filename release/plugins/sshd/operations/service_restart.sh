#!/bin/bash

SERVICE="sshd"

systemctl restart ${SERVICE}
sleep 1
systemctl status ${SERVICE}

STATUS=$(systemctl is-active ${SERVICE})
if [ "${STATUS}" != "active" ]; then
    exit -1
fi

exit 0
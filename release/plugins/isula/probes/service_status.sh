#!/bin/bash

STATUS=$(systemctl is-active isulad)
printf "isulad: [${STATUS}]"

if [ "${STATUS}" != "active" ]; then
  exit ${STATUS_CRITICAL}
fi

exit ${STATUS_HEALTHY}

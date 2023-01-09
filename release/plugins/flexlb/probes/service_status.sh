#!/bin/bash

STATUS=$(systemctl is-active flexlb)
printf "flexlb: [${STATUS}]"

if [ "${STATUS}" != "active" ]; then
  exit ${STATUS_CRITICAL}
fi

exit ${STATUS_HEALTHY}

#!/bin/bash

THRESHOLD_WARNING=80
THRESHOLD_CRITICAL=90

v_cpu_usage=$(awk '/cpu/{usage=($2+$4)*100/($2+$4+$5)} END {print usage}' /proc/stat)

python -c "
if ${v_cpu_usage} > ${THRESHOLD_CRITICAL}:
  print('cpu usage: {:.2f}% > {:.2f}%'.format(${v_cpu_usage},${THRESHOLD_CRITICAL}))
  exit(${STATUS_CRITICAL})

if ${v_cpu_usage} > ${THRESHOLD_WARNING}:
  print('cpu usage: {:.2f}% > {:.2f}%'.format(${v_cpu_usage},${THRESHOLD_WARNING}))
  exit(${STATUS_WARNING})

print('cpu usage: {:.2f}%'.format(${v_cpu_usage}))
exit(${STATUS_HEALTHY})
"
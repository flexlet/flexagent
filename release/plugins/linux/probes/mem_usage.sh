#!/bin/bash

THRESHOLD_WARNING=80
THRESHOLD_CRITICAL=90

v_mem_usage=$(free -k | awk 'NR==2 {print $3*100/$2}')

python -c "
if ${v_mem_usage} > ${THRESHOLD_CRITICAL}:
  print('mem usage: {:.2f}% > {:.2f}%'.format(${v_mem_usage},${THRESHOLD_CRITICAL}))
  exit(${STATUS_CRITICAL})

if ${v_mem_usage} > ${THRESHOLD_WARNING}:
  print('mem usage: {:.2f}% > {:.2f}%'.format(${v_mem_usage},${THRESHOLD_WARNING}))
  exit(${STATUS_WARNING})

print('mem usage: {:.2f}%'.format(${v_mem_usage}))
exit(${STATUS_HEALTHY})
"
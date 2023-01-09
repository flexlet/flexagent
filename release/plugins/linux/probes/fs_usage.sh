#!/bin/bash

# get mount points
MOUNTS=`cat /etc/fstab | grep -v '^#' | awk '{if ($2 ~ /^\//){print $2}}'`
LIST="["
for FS in ${MOUNTS}; do
  TOTAL=$(df ${FS} | grep -v '^Filesystem' | awk '{print $2}')
  USED=$(df ${FS} | grep -v '^Filesystem' | awk '{print $3}')
  LIST="${LIST}{\"FS\":\"${FS}\",\"USED\":${USED},\"TOTAL\":${TOTAL}},"
done
LIST="${LIST:0:-1}]"

# get max fs usage and output
python -c "
import json

THRESHOLD_WARNING = 80
THRESHOLD_CRITICAL = 90

LIST = json.loads('${LIST}')
MAX_USAGE = 0
OUTPUT = '%'
for ITEM in LIST:
  USAGE = 100 * ITEM['USED'] / ITEM['TOTAL']
  if USAGE > THRESHOLD_WARNING:
    STATUS = 'Critical'
  elif USAGE > THRESHOLD_CRITICAL:
    STATUS = 'Warning'
  else:
    STATUS = 'Healthy'
  
  OUTPUT = OUTPUT + ', {} => {} ({:.2f}%)'.format(ITEM['FS'], STATUS, USAGE)
  if USAGE > MAX_USAGE:
    MAX_USAGE = USAGE

print('max fs usage: {:.2f}'.format(MAX_USAGE) + OUTPUT)

if MAX_USAGE > THRESHOLD_WARNING:
  exit(${STATUS_CRITICAL})
elif MAX_USAGE > THRESHOLD_CRITICAL:
  exit(${STATUS_WARNING})
else:
  exit(${STATUS_HEALTHY})
"

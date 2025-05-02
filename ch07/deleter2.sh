#!/usr/bin/env bash

TARGET_DIR="./tmp"
LOG_FILE="deleter2.sh.log"

if [ ! -d "$TARGET_DIR" ]; then
  echo "$(date "+%Y-%m-%d %H:%M:%S") ERROR: directory '$TARGET_DIR' not found" | tee -a $LOG_FILE
  exit 1
fi

FILE_A="${TARGET_DIR}/a.txt"
FILE_B="${TARGET_DIR}/b.txt"

if [ -f "$FILE_A" ] && [ ! -f "$FILE_B" ]; then
  touch $FILE_B
  echo "$(date "+%Y-%m-%d %H:%M:%S") ERROR: created '$FILE_B'" | tee -a $LOG_FILE
  exit 1
fi

if [ ! -f "$FILE_A" ] && [ -f "$FILE_B" ]; then
  touch $FILE_A
  echo "$(date "+%Y-%m-%d %H:%M:%S") ERROR: created '$FILE_A'" | tee -a $LOG_FILE
  exit 1
fi

if [ ! -f "$FILE_A" ] && [ ! -f "$FILE_B" ]; then
  touch $FILE_A $FILE_B
  echo "$(date "+%Y-%m-%d %H:%M:%S") ERROR: created '$FILE_A' and '$FILE_B'" | tee -a $LOG_FILE
  exit 1
fi

rm $FILE_A $FILE_B
echo "$(date "+%Y-%m-%d %H:%M:%S") INFO: removed '$FILE_A' and '$FILE_B'" | tee -a $LOG_FILE
exit 0

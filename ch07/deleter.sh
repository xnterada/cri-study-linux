#!/usr/bin/env bash

TARGET_DIR="./tmp"

if [ ! -d "$TARGET_DIR" ]; then
  echo "directory '$TARGET_DIR' not found"
  exit 1
fi

FILE_A="${TARGET_DIR}/a.txt"
FILE_B="${TARGET_DIR}/b.txt"

if [ ! -f "$FILE_A" ] || [ ! -f "$FILE_B" ]; then
  touch $FILE_A $FILE_B
  exit 1
fi

rm $FILE_A $FILE_B
exit 0

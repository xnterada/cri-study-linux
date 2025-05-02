#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

MAX_FILES=9

if [ $# -ne 1 ]; then
  echo "usage: $0 <number of files>"
  exit 1
fi

if ! [[ "$1" =~ ^[0-9]+$ ]]; then
  echo "error: argument must be a number"
  exit 1
fi

if [ "$1" -lt 1 ] || [ "$1" -gt "$MAX_FILES" ]; then
  echo "error: argument should be in the range of 1-9"
  exit 1
fi

NUM_FILES=$1

for i in $(seq 1 "$NUM_FILES"); do
  FILENAME=$(printf "file%d.txt" "$i")
  touch "$FILENAME"
  echo "file '$FILENAME' created"
done

echo "all files created successfully"

exit 0
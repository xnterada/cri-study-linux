#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

if [ $UID -ne 0 ]; then
  echo "you are not root user"
  exit 1
fi

if [ -z "$1" ]; then
  echo "input username"
  echo "usage: $0 <username>"
  exit 1
fi

if ! which aws; then
  echo "AWS CLI is not installed"
  exit 1
fi

USERNAME=$1

if id -u "$USERNAME" > /dev/null 2>&1; then
  echo "user '$USERNAME' already exists"
  exit 1
fi

useradd -m $USERNAME -s /bin/bash
cd /home/$USERNAME
mkdir .ssh
chmod 700 .ssh

ssh-keygen -t rsa -b 2048 -f .ssh/id_rsa -N "" -C "$USERNAME"
cat .ssh/id_rsa.pub >> .ssh/authorized_keys && rm .ssh/id_rsa.pub
chmod 600 .ssh/authorized_keys

mv .ssh/id_rsa ".ssh/$USERNAME.pem"
aws s3 cp ".ssh/$USERNAME.pem" s3://cri-study-linux-bucket/key/

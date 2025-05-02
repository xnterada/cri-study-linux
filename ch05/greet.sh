#/usr/bin/env bash

NAME=$1

if [ -z "$NAME" ]; then
	echo "Hello, anyone?"  
else
	echo "Hello, ${NAME}!"
fi
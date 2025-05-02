#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

mv ./html/css/style.css ./html/
rm -rf ./html/css/

mv ./html/img/snake.jpg ./html/
mv ./html/img/rabbit.jpg ./html/img/rabit.jpg
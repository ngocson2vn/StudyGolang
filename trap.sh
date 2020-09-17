#!/bin/bash

set -e

trap '' ERR

export AWS_REGION=ap-northeast-1
./register-deamon-task-definition

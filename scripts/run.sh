#!/usr/bin/bash

# set root directory
cd "$(dirname "$0")"/..

# declare parameters
VERSION=1.0
OPTION=$1
ENV=$2

# import functions
. scripts/utils.sh

# switch through first arg
case $OPTION in
  build)
    docker_build $ENV
    ;;
  up)
    wrap_compose up $ENV
    ;;
  down)
    wrap_compose down $ENV
    ;;
  logs)
    docker_logs $ENV
    ;;
  env)
    cp default.env .env
    ;;
  help)
    echo "Commands: "
    echo "build <dev/prod>"
    echo "run <dev/prod>"
    echo "down <dev/prod>"
    echo "env"
    ;;
  *)
    echo "Invalid argument. Try ./run.sh help"
    ;;
esac
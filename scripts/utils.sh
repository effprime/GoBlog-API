docker_build() {
  if [[ $1 == "prod" ]]
  then
    tag=$VERSION
  else
    tag=dev-$VERSION
  fi
  docker build -t effprime/goblog-api:$tag .
}
docker_logs() {
  if [[ $1 == "prod" ]]
  then
    container=goblog_api_prod
  elif [[ $1 == "dev" ]]
  then
    container=goblog_api_dev
  fi
  docker logs $container -f
}
wrap_compose() {
  prod_compose="docker-compose"
  dev_compose="$prod_compose -f docker-compose.postgres.yaml -f docker-compose.yaml -f docker-compose.override.yaml"
  if [[ $1 == "up" ]]
  then
    if [[ $2 == "prod" ]]
    then
      $prod_compose up -d
    elif [[ $2 == "dev" ]]
    then
      $dev_compose up -d
    else
      echo "Invalid input for env setting: dev/prod"
    fi
  elif [[ $1 == "down" ]]
  then
    if [[ $2 == "prod" ]]
    then
      $prod_compose down
    elif [[ $2 == "dev" ]]
    then
      $dev_compose down
    else
      echo "Invalid input for env setting: dev/prod"
    fi
  fi
}
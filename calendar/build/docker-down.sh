#!/bin/bash
export TZ=Europe/Samara
export POSTGRES_USER=postgres
export POSTGRES_DB=pg_calendar
export POSTGRES_PASSWORD=12345

docker-compose  --file build/deploy/docker-compose.yml down 

docker image rm calendar_sender --force
docker image rm calendar_scheduler --force
docker image rm calendar_api --force
docker image rm calendar_web --force
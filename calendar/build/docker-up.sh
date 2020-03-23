#!/bin/bash

export TZ=Europe/Samara
export POSTGRES_USER=postgres
export POSTGRES_DB=pg_calendar
export POSTGRES_PASSWORD=12345

#docker build -t calendar_db:latest -f ./build/docker/db/Dockerfile .

#docker-compose -f build/deploy/docker-compose.yml up --detach #--remove-orphans
docker-compose -f build/deploy/docker-compose.yml up --build --detach #--remove-orphans

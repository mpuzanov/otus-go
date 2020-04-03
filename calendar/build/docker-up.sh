#!/bin/bash

export TZ=Europe/Samara
export POSTGRES_USER=postgres
export POSTGRES_DB=pg_calendar_test
export POSTGRES_PASSWORD=12345

#docker build -t calendar_db:latest -f ./build/db/Dockerfile .
#docker build -t calendar_db_test:latest -f ./build/db-test/Dockerfile .

docker-compose -f build/docker-compose.yml up --build --detach #--remove-orphans
#docker-compose -f build/docker-compose2.yml up --detach

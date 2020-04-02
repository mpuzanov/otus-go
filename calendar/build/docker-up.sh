#!/bin/bash

export TZ=Europe/Samara

#docker build -t calendar_db:latest -f ./build/db/Dockerfile .
#docker build -t calendar_db_test:latest -f ./build/db-test/Dockerfile .

#docker-compose -f build/docker-compose.yml up --detach

docker-compose -f build/docker-compose.yml up --build --detach #--remove-orphans
#docker-compose -f build/docker-compose2.yml up --build --detach

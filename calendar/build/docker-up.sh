#!/bin/bash

export TZ=Europe/Samara
export POSTGRES_USER=postgres
export POSTGRES_DB=pg_calendar
export POSTGRES_PASSWORD=12345

docker-compose -f build/deploy/docker-compose.yml up --build --detach #--remove-orphans

#!/bin/bash
docker-compose  --file build/docker-compose.yml down #--remove-orphans

#docker image rm calendar_db:latest --force
docker image rm calendar_db --force
docker image rm calendar_db_test --force
docker image rm calendar_sender --force
docker image rm calendar_scheduler --force
docker image rm calendar_api --force
docker image rm calendar_web --force
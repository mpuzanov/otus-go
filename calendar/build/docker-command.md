# rabbitmq	
docker run -d --name rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:3-management

## Перезапускаем контейнер
docker stop rabbitmq
docker start rabbitmq

## админка
http://localhost:15672
login: guest pswd: guest


docker build -t puzanovma/calendar -f ./deployments/calendar/Dockerfile .
docker run --rm -it -p 50051:50051 -e TZ=Europe/Samara -e DB_URL=postgres://postgres:12345@0.0.0.0:5432/pg_calendar?sslmode=disable --name calendar puzanovma/calendar
docker run -d -p 50051:50051 -e TZ=Europe/Samara --name calendar puzanovma/calendar

## Удаление всех остановленных контейнеров
    docker rm $(docker ps -a -q -f status=exited)
## Удаление недействительных образов
    docker rmi $(docker images -q -f dangling=true)
## Удаление недействительных томов
	docker volume ls -f dangling=true
	docker volume rm $(docker volume ls -f dangling=true -q)
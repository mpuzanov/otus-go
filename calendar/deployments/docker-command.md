# rabbitmq	
docker run -d --name rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:3-management

## Перезапускаем контейнер
docker stop rabbitmq
docker start rabbitmq

## админка
http://localhost:15672
login: guest pswd: guest


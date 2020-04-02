FROM rabbitmq:3.8-management-alpine 

RUN apk add --no-cache curl

EXPOSE 4369 5671 5672 25672 15671 15672
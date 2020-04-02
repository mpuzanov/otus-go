FROM postgres:12.0-alpine

COPY ./build/db/initdb.sh /docker-entrypoint-initdb.d/initdb.sh
RUN chmod -R 755 /docker-entrypoint-initdb.d/
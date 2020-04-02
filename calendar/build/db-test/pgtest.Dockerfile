FROM postgres:12.0-alpine

COPY ./build/db-test/db-test.sql /docker-entrypoint-initdb.d/db-test.sql
RUN chmod -R 755 /docker-entrypoint-initdb.d/


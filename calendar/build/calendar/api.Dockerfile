FROM alpine:3.11
ENV APP_NAME calendar_api
LABEL name=${APP_NAME} maintainer="Mikhail Puzanov <mpuzanov@mail.ru>" version="1"

WORKDIR /opt/${APP_NAME}
COPY ./calendar_api ./bin/
COPY ./configs/prod/config.yaml ./configs/
RUN apk add --no-cache tzdata \
    && apk add -U --no-cache ca-certificates \
    && adduser -D -g appuser appuser \
    && chmod -R 755 ./

USER appuser

ENTRYPOINT ["./bin/calendar_api"]
CMD ["grpc_server", "-c", "./configs/config.yaml"]
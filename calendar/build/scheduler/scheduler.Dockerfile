FROM alpine:3.11
ENV APP_NAME scheduler
LABEL name=${APP_NAME} maintainer="Mikhail Puzanov <mpuzanov@mail.ru>" version="1"

WORKDIR /opt/${APP_NAME}
COPY ./calendar_scheduler ./bin/
COPY ./configs/prod/config-scheduler.yaml ./configs/
RUN apk add --no-cache tzdata \
    && apk add -U --no-cache ca-certificates \
    && adduser -D -g appuser appuser \
    && chmod -R 755 ./

USER appuser

ENTRYPOINT ["./bin/calendar_scheduler"]
CMD ["-c", "./configs/config-scheduler.yaml"]


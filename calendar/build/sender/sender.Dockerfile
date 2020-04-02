FROM alpine:3.11
ENV APP_NAME calendar_sender
LABEL name=${APP_NAME} maintainer="Mikhail Puzanov <mpuzanov@mail.ru>" version="1"

WORKDIR /opt/${APP_NAME}
COPY ./calendar_sender ./bin/
COPY ./configs/prod/config-sender.yaml ./configs/
 
RUN chmod -R 755 ./ \
    && adduser -D -g appuser appuser \
    && apk add --no-cache tzdata \
    && apk add -U --no-cache ca-certificates

USER appuser

ENTRYPOINT ["./bin/calendar_sender"]
CMD ["./bin/calendar_sender",  "-c", "./configs/config-sender.yaml"]
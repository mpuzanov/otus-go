FROM golang:1.13-alpine as builder
ENV APP_NAME scheduler
RUN apk add --no-cache tzdata \
    && apk add -U --no-cache ca-certificates \
    && adduser -D -g appuser appuser
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${APP_NAME} ./cmd/${APP_NAME}

FROM scratch
ENV APP_NAME scheduler
LABEL name=${APP_NAME} maintainer="Mikhail Puzanov <mpuzanov@mail.ru>" version="1"
WORKDIR /opt/${APP_NAME}
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /opt/${APP_NAME}/bin/${APP_NAME} ./bin/
COPY --from=builder /opt/${APP_NAME}/configs/config-prod.yaml ./configs/
EXPOSE 8888
ENTRYPOINT ["./bin/scheduler"]
CMD ["-c", "./configs/config-prod.yaml"]

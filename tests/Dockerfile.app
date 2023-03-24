ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS builder

ARG APP_TYPE

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
COPY . /pubsubmit
WORKDIR /pubsubmit/tests/integration-${APP_TYPE}
RUN go mod download
RUN go build -o ./app ./main.go

FROM alpine:latest

ARG APP_TYPE

WORKDIR /root/
COPY --from=builder ./pubsubmit/tests/integration-${APP_TYPE}/app ./
EXPOSE 8080
ENTRYPOINT ["./app"]
FROM golang:alpine as base

WORKDIR /src

RUN apk update && \
    apk add --no-cache git curl openssl build-base ca-certificates 

ADD . /src

RUN go build -o app

FROM alpine

RUN apk update && \
    apk add --no-cache curl openssl ca-certificates &&\
    mkdir -p /usr/local/logs

COPY --from=base /src/app /usr/local/bin

ENV http_proxy ''
ENV https_proxy ''

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/app"]
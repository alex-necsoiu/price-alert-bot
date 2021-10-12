FROM golang:latest

RUN mkdir/build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get   git clone github.com/alex-necsoiu/uphold-bot

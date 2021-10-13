FROM golang:latest

RUN mkdir ./build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get -v -u github.com/alex-necsoiu/uphold-bot/blob/c6b4f4b6bcf1295eeb0a0a958bc869fdaaa5442c/bot/pkg/build/main
RUN cd /build/main && git clone https://github.com/alex-necsoiu/uphold-bot.git
RUN cd /build/main && go build 

ENTRYPOINT ["/build/main/main"]
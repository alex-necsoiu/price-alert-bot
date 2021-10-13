FROM golang:latest

RUN mkdir ./build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get   git clone https://github.com/alex-necsoiu/uphold-bot/tree/master/bot/pkg/build/main
RUN cd /build/main && git clone https://github.com/alex-necsoiu/uphold-bot.git
RUN cd /build/main && go build 

ENTRYPOINT ["/build/main/main"]
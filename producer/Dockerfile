FROM golang:1.13.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

RUN go get github.com/Shopify/sarama.v1.20.1

ENV SOURCES /go/src/github.com/PacktPublishing/Advanced-Cloud-Native-Go/Communication/Kafka/
COPY . ${SOURCES}

RUN cd ${SOURCES}producer/ && CGO_ENABLED=0 go build

ENV BROKER_ADDR localhost:9092

WORKDIR ${SOURCES}producer/
CMD ${SOURCES}producer/producer

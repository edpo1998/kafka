FROM golang:1.13.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

ENV SOURCES /go/src/github.com/PacktPublishing/Advanced-Cloud-Native-Go/Communication/Kafka/

COPY . ${SOURCES}

RUN cd ${SOURCES}subscriber/ && CGO_ENABLED=0 go build

ENV BROKER_ADDR localhost:9092

WORKDIR ${SOURCES}subscriber/
CMD ${SOURCES}subscriber/subscriber

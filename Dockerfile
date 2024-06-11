FROM golang:1.21

ENV GO111MODULE=on

WORKDIR /app

COPY . /app

RUN go run main.go
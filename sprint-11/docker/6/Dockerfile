FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN go build -o main

CMD ./main

FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN go build -o main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ./main
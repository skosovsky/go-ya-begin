FROM golang:1.22.1

WORKDIR /app

COPY go.mod /

RUN go mod download

COPY . .

RUN go build -o main

CMD ./main
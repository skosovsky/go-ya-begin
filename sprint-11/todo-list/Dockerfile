FROM golang:latest
LABEL authors="skosovsky"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./todo_list
EXPOSE 8080
CMD ["./todo_list"]
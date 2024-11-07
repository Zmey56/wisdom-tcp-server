FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go build -o client cmd/client/main.go
EXPOSE 8080

CMD ["./client"]

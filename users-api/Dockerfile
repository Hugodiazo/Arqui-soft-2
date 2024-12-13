# Dockerfile para users-api
FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o users-api

EXPOSE 8080

CMD ["./users-api"]
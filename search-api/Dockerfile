# Dockerfile para search-api
FROM golang:1.22

WORKDIR /app

COPY ../../wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

COPY . .

RUN go mod tidy
RUN go build -o search-api

EXPOSE 8080

CMD ["/wait-for-it.sh", "rabbitmq:5672", "--", "./search-api"]
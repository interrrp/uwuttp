FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

ENV ADDR=:80
ENV REDIS_ADDR=cache:6379
EXPOSE 80

CMD ["./main"]

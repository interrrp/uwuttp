FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -ldflags "-s -w" -o app ./cmd

ENV ADDR=:80
EXPOSE 80

CMD ["./app"]

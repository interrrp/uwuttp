FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags "-s -w" -o main .

ENV ADDR=:80
EXPOSE 80

CMD ["./main"]

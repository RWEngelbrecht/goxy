FROM golang:1.16-alpine

WORKDIR /app
COPY . /app

CMD go run .

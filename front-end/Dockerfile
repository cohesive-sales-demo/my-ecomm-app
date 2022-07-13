FROM golang:1.17-alpine

RUN apk add build-base

RUN mkdir -p /app
WORKDIR /app
COPY . /app

RUN go build ./front-end/main.go
EXPOSE 8080

CMD ./main
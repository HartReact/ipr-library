FROM golang:1.20.3-alpine3.17

COPY ./app ./app

WORKDIR ./app

RUN go build -o app ./cmd/main.go
EXPOSE 8080
CMD ["./app"]
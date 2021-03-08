FROM golang:1.16 AS build_base

RUN mkdir /app
WORKDIR /app

ENV GOPROXY "https://proxy.golang.org"

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o bin/main main.go

CMD ["/app/bin/main"]

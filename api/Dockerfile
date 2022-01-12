FROM golang:1.16 AS build
# setup working directory
RUN mkdir /app
WORKDIR /app
# make sure proxy is correct
ENV GOPROXY "https://proxy.golang.org"
# copy mod files
COPY go.mod .
COPY go.sum .
# install mods
RUN go mod download
RUN go mod verify
# copy code
COPY . .
# build binary
RUN CGO_ENABLED=0 go build -o bin/main main.go

# thin running layer
FROM alpine:3.15.0 AS run
WORKDIR /app
COPY --from=build /app/bin/main .
CMD ["./main"]

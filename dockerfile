FROM golang:alpine as builder

WORKDIR /go/src/Noteus
COPY . .

RUN go env -w GO111MODULE=on \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build .

FROM alpine:latest

WORKDIR /go/src/Noteus

ENTRYPOINT ["./main"]

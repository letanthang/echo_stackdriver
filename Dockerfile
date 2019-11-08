FROM golang:1.13.3-alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .

#download dependencies
RUN go mod download

#build
RUN go build -o /go/bin/main

FROM alpine:3.5
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

WORKDIR /app
COPY ./config/config.yaml ./config/

COPY --from=builder go/bin/main .
EXPOSE 9090
ENTRYPOINT ["./main"]

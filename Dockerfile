FROM golang:1.15 as builder
COPY . /
WORKDIR /
RUN go build

FROM debian:stable-slim
COPY --from=builder /ping-pong /usr/bin/
ENTRYPOINT ["ping-pong"]
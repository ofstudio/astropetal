FROM golang:1.16-alpine AS builder
WORKDIR /src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /build/astropetal-bot ./cmd

FROM alpine:latest
COPY --from=builder /build/astropetal-bot /
CMD ["/astropetal-bot"]

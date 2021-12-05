FROM golang:1.17-alpine AS builder
WORKDIR /src
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /build/astropetal-bot ./cmd/astropetal-bot.go

FROM alpine:latest
COPY --from=builder /build/astropetal-bot /
USER nobody
CMD ["/astropetal-bot"]

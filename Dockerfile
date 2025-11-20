FROM golang:1.24-bullseye AS builder
ENV GONOPROXY=

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/interfacer-dpp ./cmd/main

FROM ghcr.io/dyne/zenroom:latest

WORKDIR /root

ENV HOST=0.0.0.0
ENV PORT=8080
ENV GIN_MODE=release

COPY --from=builder /usr/local/bin/interfacer-dpp /usr/local/bin/interfacer-dpp

EXPOSE 8080

# IMPORTANT: override both entrypoint and cmd so zenroom doesn't run first
ENTRYPOINT ["/usr/local/bin/interfacer-dpp"]
CMD []
FROM golang:1.24-bullseye AS builder
ENV GONOPROXY=

RUN wget -O /usr/local/zenroom-zencode-exec \
    https://github.com/dyne/zenroom/releases/latest/download/zencode-exec \
 && chmod +x /usr/local/zenroom-zencode-exec

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /interfacer-dpp ./cmd/main

FROM dyne/devuan:chimaera
WORKDIR /root

ENV HOST=0.0.0.0
ENV PORT=8080
ENV GIN_MODE=release
ENV PATH="/usr/local/zenroom/bin:${PATH}"

RUN mkdir -p /usr/local/zenroom/bin

EXPOSE 8080

COPY --from=builder /interfacer-dpp /usr/local/bin/interfacer-dpp
COPY --from=builder /usr/local/zenroom-zencode-exec /usr/local/zenroom/bin/zencode-exec

CMD ["/usr/local/bin/interfacer-dpp"]
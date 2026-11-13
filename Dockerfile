FROM golang:1.24-bookworm AS builder
ENV GONOPROXY=

# Install build dependencies for zencode-exec
RUN apt-get update && apt-get install -y build-essential cmake git

# Download or build zencode-exec based on architecture
RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        wget -O /usr/local/zenroom-zencode-exec \
            "https://github.com/dyne/zenroom/releases/latest/download/zencode-exec" && \
        chmod +x /usr/local/zenroom-zencode-exec; \
    else \
        echo "Building zencode-exec from source for $ARCH..." && \
        git clone --depth 1 https://github.com/dyne/Zenroom.git /tmp/zenroom && \
        cd /tmp/zenroom && \
        make linux-zencode-exec && \
        cp zencode-exec /usr/local/zenroom-zencode-exec && \
        chmod +x /usr/local/zenroom-zencode-exec; \
    fi

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /interfacer-dpp ./cmd/main

# Devuan 5 (Debian 12 base): the same glibc the builder linked against.
FROM dyne/devuan:daedalus
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*
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
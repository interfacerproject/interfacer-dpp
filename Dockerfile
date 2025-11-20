FROM golang:1.22-bullseye AS builder
ENV GONOPROXY=

RUN apt-get update \
	&& apt-get install -y --no-install-recommends build-essential git cmake vim python3 python3-pip zsh libssl-dev ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

RUN pip3 install --no-cache-dir meson ninja \
	&& git clone https://github.com/dyne/Zenroom.git /zenroom

RUN cd /zenroom && make linux-go

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
COPY --from=builder /zenroom/build/src/zenroom-exec /usr/local/zenroom/bin/zenroom-exec
COPY --from=builder /zenroom/meson/libzenroom.so /usr/lib/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libssl.so.1.1 /lib/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libcrypto.so.1.1 /lib/

CMD [ "interfacer-dpp" ]

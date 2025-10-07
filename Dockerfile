FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /interfacer-dpp ./cmd/main

EXPOSE 8080

CMD [ "/interfacer-dpp" ]

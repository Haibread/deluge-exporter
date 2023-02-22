FROM golang:1.20-alpine as builder

RUN apk upgrade --update-cache --available
RUN apk add --no-cache \
        gcc \
        musl-dev 

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o deluge-exporter ./main.go

FROM alpine:3.17.2

RUN apk --no-cache add ca-certificates

LABEL org.opencontainers.image.source="https://github.com/Haibread/deluge-exporter"

VOLUME [ "/app/config" ]
WORKDIR /app

COPY --from=builder /app/deluge-exporter /app/

ENTRYPOINT ["/app/deluge-exporter"]
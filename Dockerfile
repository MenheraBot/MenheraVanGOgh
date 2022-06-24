FROM golang:1.18-stretch AS builder

ENV USER=appuser
ENV UID=10001

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN go build -ldflags="-s -w" -o vangogh src/main.go

FROM gcr.io/distroless/base
WORKDIR app

COPY --from=builder /app/vangogh vangogh
COPY assets assets

USER root
EXPOSE 2080

ENTRYPOINT ["/app/vangogh"]
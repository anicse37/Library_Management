# -------- 1. Builder stage --------
FROM golang:1.23.10-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/server

RUN go build -o /app/library-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/library-app .

COPY --from=builder /app/internal/template ./internal/template

EXPOSE 5050

CMD ["./library-app"]

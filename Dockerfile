# Builder stage
FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app/

COPY ./ ./

RUN go build -o ./main ./cmd/bookkeeper/main.go

# Runner stage
FROM alpine:3.18

WORKDIR /app/

COPY --from=builder /app/main ./main
COPY ./scripts/ ./scripts/
COPY ./configs/config.env ./configs/config.env
COPY ./internal/db/migrations/ ./internal/db/migrations/

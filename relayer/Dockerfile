# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o relayer-bin ./cmd/relayer

# Runtime stage  
FROM alpine:latest

# Install required packages
RUN apk --no-cache add \
    ca-certificates \
    postgresql-client \
    wget

# Install golang-migrate
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/ && \
    rm migrate.linux-amd64.tar.gz

WORKDIR /app

COPY --from=builder /app/relayer-bin .
COPY --from=builder /app/.env.staging.yaml .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/contracts/out ./contracts/out
COPY --from=builder /app/entrypoint.sh .

RUN chmod +x ./relayer-bin ./entrypoint.sh

# ОБНОВЛЕНО: Добавляем порт для метрик Prometheus
EXPOSE 8080 9090

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

ENTRYPOINT ["./entrypoint.sh"]

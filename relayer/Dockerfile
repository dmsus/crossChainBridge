# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o relayer-bin ./cmd/relayer

# Runtime stage  
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/relayer-bin .
COPY --from=builder /app/.env.staging.yaml .
# Копируем ABI файлы из контекста сборки
COPY --from=builder /app/contracts/out ./contracts/out

RUN chmod +x ./relayer-bin

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./relayer-bin", "--env", "staging"]

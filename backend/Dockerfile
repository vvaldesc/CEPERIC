# Multi-stage build para producción

# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar dependencias
COPY go.mod go.sum* ./
RUN go mod download

# Copiar código fuente
COPY . .

# Build optimizado
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main ./cmd/api

# Stage 2: Runtime
FROM alpine:latest

# Instalar certificados CA para HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar binario compilado
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8080

# Usuario no-root para seguridad
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /root

USER appuser

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/v1/health || exit 1

# Ejecutar
CMD ["./main"]

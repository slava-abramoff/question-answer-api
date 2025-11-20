# Stage 1: Build the Go binary
FROM golang:1.25 AS builder

WORKDIR /app

# Копируем go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники и собираем приложение
COPY ./application ./application
RUN go build -o app ./application/cmd/main.go

# Stage 2: Runtime image
FROM debian:bookworm-slim

WORKDIR /app

# Устанавливаем минимальные утилиты для работы приложения
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Копируем собранное приложение из builder stage
COPY --from=builder /app/app .

# По умолчанию запускаем приложение
CMD ["./app"]

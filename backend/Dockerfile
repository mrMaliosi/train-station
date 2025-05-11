# syntax=docker/dockerfile:1

FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main-exe ./cmd/main.go

# --- Релизный контейнер ---
FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main-exe /app/

# Копируем файл секретов
COPY .env .

# Экспортируем порт
EXPOSE 8080

CMD ["./main-exe"]

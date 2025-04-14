FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

# Копируем бинарный файл и .env из предыдущего этапа
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

CMD ["./main"]
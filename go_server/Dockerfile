# go_server/Dockerfile
FROM golang:1.24-alpine AS builder

# Устанавливаем переменные для Go proxy
ENV GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV CGO_ENABLED=0

WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY go.mod ./
COPY go.sum ./

# Копируем весь исходный код go_server в правильную структуру
COPY go_server/ ./go_server/

# Загружаем зависимости
RUN go mod download

# Собираем бинарник, указывая правильный путь
RUN go build -o main ./go_server/cmd/main.go

# Финальная стадия
FROM alpine:latest

# Устанавливаем curl для healthcheck
RUN apk add --no-cache curl ca-certificates

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/main .

# Сохраняем структуру папок для совместимости с кодом
COPY go_server/ ./go_server/

# Также копируем в корень для удобства
COPY go_server/static/ ./static/
COPY go_server/templates/ ./templates/

# Копируем .env файл
COPY .env ./

EXPOSE 8000

CMD ["./main"]
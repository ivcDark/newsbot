# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

# Установка зависимостей
RUN apk add --no-cache bash curl tzdata cron

WORKDIR /app

# Копируем go.mod/sum и качаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники и собираем
COPY . .
RUN go build -o news ./main.go

# Копируем кронтаб в системный файл пользователя root
COPY ./_docker/Cron/newsbot-cron /etc/crontabs/root

# Убедимся, что скрипт исполняемый
RUN chmod +x /app/news

# Cron запускается в foreground (без этого контейнер завершится)
CMD ["crond", "-f", "-l", "2"]

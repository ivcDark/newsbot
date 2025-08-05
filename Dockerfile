# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

# Установка tini, cron и других утилит
RUN apk add --no-cache bash curl tzdata dcron tini

# Установка tini как init-процесс
ENTRYPOINT ["/sbin/tini", "--"]

WORKDIR /app

# go mod deps
COPY go.mod go.sum ./
RUN go mod download

# код и сборка
COPY . .
RUN go build -o /usr/local/bin/newsbot ./main.go

# Кронтаб
COPY ./newsbot-cron.txt /etc/crontabs/root

# Разрешаем выполнение бинарника
RUN chmod +x /usr/local/bin/newsbot

# Запуск cron
CMD ["crond", "-f", "-l", "8"]

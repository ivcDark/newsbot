FROM newsbot-app:latest AS cron-build

FROM alpine:3.18

COPY --from=app-build /app/newsbot /app/newsbot

# Установки для cron
RUN apk add --no-cache busybox-cron

COPY newsbot-cron /etc/crontabs/root

CMD ["crond", "-f", "-d", "8"]
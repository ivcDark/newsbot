#!/bin/sh

# Если передана команда — запускаем её
if [ $# -gt 0 ]; then
  exec ./newsbot "$@"
else
  echo "Запускаем cron-сервис..."
  crond -f -l 2
fi

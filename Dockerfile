FROM golang:1.23 AS builder

WORKDIR /app

# Копируем исходные файлы Go в контейнер
COPY . /app

# Загружаем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o url-shortener .

FROM debian:bullseye-slim

WORKDIR /root

# Копируем собранный бинарник из предыдущего этапа
COPY --from=builder /app/url-shortener .

# Открываем нужный порт
EXPOSE 8080

CMD ["./url-shortener"]

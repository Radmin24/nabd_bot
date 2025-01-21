FROM golang:1.23 AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем сначала только go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Теперь копируем весь исходный код
COPY . .

# Сборка Go-программы
RUN CGO_ENABLED=0 GOOS=linux go build -o nbd_bot ./main.go 

# Используем Debian для минимального размера
FROM debian:bullseye-slim

# Копируем скомпилированный Go-бинарник из предыдущего этапа
COPY --from=build /app/nbd_bot /usr/local/bin/nbd_bot

# Изменяем права доступа для nbd_bot и выводим содержимое /usr/local/bin/
RUN chmod +x /usr/local/bin/nbd_bot \
    && ls -l /usr/local/bin/

RUN apt-get update && apt-get install -y ca-certificates curl

# Монтируем сокет Docker для взаимодействия с контейнерами
VOLUME ["/var/run/docker.sock:/var/run/docker.sock"]

# Команда запуска вашего бота
CMD ["/usr/local/bin/nbd_bot"]
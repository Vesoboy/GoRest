# Используем официальный образ Golang
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /main

# Копируем go.mod и go.sum для оптимизации кеширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем весь проект в контейнер
COPY . .

# Собираем приложение
RUN go build -o /app/app

# Указываем команду для запуска
CMD ["/app/app"]

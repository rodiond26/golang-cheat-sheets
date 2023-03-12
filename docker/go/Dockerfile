# Взять базовый образ
FROM golang as builder

# Создать директорию внутри котейнера /app
RUN mkdir /app

# Скопировать все внуть котейнера в /app
COPY . /app

# Все следующие комнады будут выполняться относительно диресктории workdir=/app
WORKDIR /app

# Выполнить команду внутри контейнера после старта
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# Взять пустой образ специальный пустой образ scratch (0 байт)
FROM scratch

# Скопировать испролняемый файл server в корень образа scratch
COPY --from=builder /app/server /

# Приложение слушает порт 8080
EXPOSE 8080

# Запуск бинарного файла server при запуске контейнера
CMD ["/server"]

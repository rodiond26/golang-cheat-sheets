Собрать образ
```console
docker build . -t myapp
```

Запустить контейнер
```console
docker run -d -p 8080:8080 myapp
```

Проверить приложение

http://localhost:8080/api/v1/sum?a=101&b=202

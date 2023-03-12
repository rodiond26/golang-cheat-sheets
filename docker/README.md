### Работа с Docker

---

##### Остановить все контейнеры (git bash):
```console
docker stop $(docker ps -q)
```

##### Очистить все неиспользуемые или не связанные с контейнерами образы, контейнеры, тома и сети
```console
docker system prune -a
```
##### Очистить все тома
```console
docker volume prune
```


Запустить контейнер

Остановить контейнер по имени

`docker stop container_name`

Посмотреть список контейнеров

Подключиться к терминалу внутри контейнера

Собрать приложение go в одном контейнере и запустить бинарный файл в другом контейнере


https://youtu.be/lmTpT4eDZQk?t=4672

https://github.com/jm-program/docker-kubernetes-workshop/blob/master/1.docker/docker.pdf

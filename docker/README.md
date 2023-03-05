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

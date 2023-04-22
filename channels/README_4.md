### Select

---

#### Чтение из нескольких каналов

```
// Объявить каналы
errors := make(chan error)
cancelled := make(chan string)

// Зарегистрировать как слушатели
ch.NotifyCancel(cancelled)
ch.NotifyClose(errors)

// Чтение из нескольких каналов
for {
    select {
        case err := <- errors:
            log.printf("error received : %s", err)
            panic(err)
        case msg := <- cancelled:
            log.printf("message cancelled")
    }
}
```

---

#### Таймаут ожидания завершения операции

```
// Вызвать функцию в горутине и записать ее результат в канал
resultChan := make(chan int)

go func () {
    resultChan <= veryExpensiveOperation()
}

// Вызвать таймер, если не пришел результат функции за заданное время
timer := time.NewTimer(1 * time.Second)

select {
    case value := <- resultChan:
        log.printf("got result")
    case <- timer.C:
        return errors.New("timeout")
    }
}
```

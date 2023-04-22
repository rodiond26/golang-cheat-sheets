### Mutex

---

#### Стандартное использование mutex

```
// Создать структуру
type SomeStruct struct {
    values map[string]int
    mutex sync.Mutex
}

// Использовать mutex
func (s *SomeStruct) Count(key string) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    value, ok := s.values[key]
    if !ok {
        s.values[key] = 1
    } else {
        s.values[key] = value + 1
    }
}
```

#### Ожидание окончания всех горутин с помощью mutex

```
// Создать переменные
workerCount := 24
waitGroup := sync.WaitGroup{}

// Запустить все горутины
for i := 0; i < workerCount; i++ {
    waitGroup.Add(1)
    go func() {
        someFunction()
        // ...
        waitGroup.Done()
    }()
}

// Ожидание завершения всех горутин
waitGroup.Wait()
```

### Once

---

#### Запуск части кода только один раз

```
// Создать структуру
type SomeStruct struct {
    once sync.Once
    mutex sync.Mutex
    data *list.List
}

// Запустить функцию
func (s *SomeStruct) DoSomething(e string) {
    s.onceDo(func() {
        s.data = list.New()
    }()
    
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.data.PushBack(e)
}
```

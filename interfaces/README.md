### Использование интерфейсов

---

* [Простой интерфейс](./simple/main.go)
* [Пример с http](./http/main.go)
* [Пример с Writer](./custom/main.go)


* Интерфейс - это контракт поведения других тип. Интерфейс показывает поведение и скрывает реализацию.
* Интрерфейс помогают реализовать принцип полиморфизма из ООП - возможность передавать в функцию аргументы различных
  типов.
* Объявление интерфейса:

```go
type Writer interface {
Write(p []byte)(n int, err error)
}
```

* Пример интерфейса в коде.

* Утверждение типа (type assertion):

```go
phone, ok := s.(*Phone) // Утверждение типа
if ok {
fmt.Println("Balance", phone.Balance)
}
```

* Выбор типа (type switch):

```go
switch s.(type) { // type switch - выбор типа
case *Email:
fmt.Println("switch case: *Email")
case *Phone:
fmt.Println("switch case: *Phone")
phone := s.(*Phone)
fmt.Println("Balance", phone.Balance)
}
```

* Встраивание интерфейсов

```go
type Sender interface {
Send(msg string) error
}

type Caller interface {
Call(number int) error
}

type IPhone interface {
Caller
Sender
}
```

* Пустой интерфейс - это тип интерфейса, к которому нет никаких требований к реализации, т. е. не нужно реализовывать
  никакие методы, чтобы соответствовать этому интерфейсу, т. к. он пустой.

```go
i := inteface{}
```

* [Пример путого интерфейса](./empty_interface/main.go)

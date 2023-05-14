package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Address string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено по почте на адрес %v \n", msg, e.Address)
	return nil
}

// Пустой интерфейс
func Notify(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("switch case: int")
	}
	s, ok := i.(Sender)
	if !ok {
		fmt.Println("Ошибка")
		return
	}

	err := s.Send("Сообщение для пустого интерфейса")
	if err != nil {
		fmt.Println("Ошибка для пустого интерфейса")
	}
	fmt.Printf("type = %T value = %v\n", i, i)
}

func main() {
	Notify(1)
	Notify("message")
	var s struct{}
	Notify(s)
	email := &Email{"email"}
	Notify(email)
}

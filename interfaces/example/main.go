package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Caller interface {
	Call(number int) error
}

// IPhone чтобы соответствовать интерфейсу IPhone,
// необходимо реализовать все методы интерфейсов Caller и Sender и функции MyFunc() и MyFunc2()
type IPhone interface {
	Caller
	Sender
	MyFunc()
	MyFunc2()
}

type Email struct {
	Address string
}

type Phone struct {
	Number  int
	Balance int
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено по телефону с номера %v \n", msg, p.Number)
	return nil
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено по почте на адрес %v \n", msg, e.Address)
	return nil
}

// Notify функция принимает аргумент типа Sender
func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch s.(type) { // type switch - выбор типа
	case *Email:
		fmt.Println("switch case: *Email")
	case *Phone:
		fmt.Println("switch case: *Phone")
		phone := s.(*Phone)
		fmt.Println("Balance", phone.Balance)
	}

	phone, ok := s.(*Phone) // type assertion - утверждение типа
	if ok {
		fmt.Println("Balance", phone.Balance)
	}

	fmt.Println("Success")
}

func main() {
	email := &Email{"dev@mail.com"}
	Notify(email)

	phone := &Phone{Number: 777, Balance: 100}
	Notify(phone)
}

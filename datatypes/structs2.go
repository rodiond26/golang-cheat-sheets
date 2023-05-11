package main

import "fmt"

type Client struct {
	ID   int64
	Age  int
	Name string
	Person
}

func (c Client) GetName() string {
	return c.Name
}

type Person struct {
	Name   string
	ID     int64
	Status string
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetStatus() string {
	return p.Status
}

func main() {
	c1 := Client{}
	c2 := Client{}

	c1.Person.Name = "Person Name 1" // Обычный доступ к полю
	c2.Name = "Person Name 2"        // Доступ к встраиваемой структуре

	c1.Person.ID = 1 // Обычный доступ к полю
	c2.ID = 2        // Доступ к полю самого верхнего уровня
	c1.Person.Status = "Person Status"

	fmt.Printf("c1 = %#v\n", c1)
	fmt.Printf("c2 = %#v\n", c2)
	fmt.Println(c1.GetName())        // Можно вызвать метод структуры Client
	fmt.Println(c1.Person.GetName()) // Можно вызвать метод структуры Person
	fmt.Println(c1.GetStatus())      // Если у структуры Client нет метода GetStatus(), то метод GetStatus() расширится из структуры Person
}

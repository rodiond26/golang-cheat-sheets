package main

import "fmt"

//Решите задачу:
//Нам необходимо реализовать фигуры. Базово, предположим у нас есть несколько
//фигур с различными свойствами. Давайте возьмем Rectangle и Circle. Над каждой
//из таких фигур можно выполнять следующие действия:
//Area() - возвращает в результате площадь фигуры
//Type() - возвращает строковое название (окружность, прямоугольник)
//Реализуйте такой набор фигур с выводом на экран значений их действий.

func main() {
	r := NewRectangle("rec 1", 3.25, 5.25)
	c := NewCircle("cir 1", 3.33)

	a1 := r.Area()
	a2 := c.Area()

	t1 := r.Type()
	t2 := c.Type()

	fmt.Printf("type: %v area: %v\n", t1, a1)
	fmt.Printf("type: %v area: %v\n", t2, a2)
}

const (
	PI float64 = 3.14
)

type Rectangle struct {
	name  string
	sideA float64
	sideB float64
}

type Circle struct {
	name   string
	radius float64
}

type Figurer interface {
	Areaer
	Typer
}

type Areaer interface {
	Area() float64
}

type Typer interface {
	Type() string
}

func NewRectangle(name string, sideA, sideB float64) *Rectangle {
	return &Rectangle{
		name:  name,
		sideA: sideA,
		sideB: sideB,
	}
}

func NewCircle(name string, radius float64) *Circle {
	return &Circle{
		name:   name,
		radius: radius,
	}
}

func (r *Rectangle) Area() float64 {
	return r.sideA * r.sideB
}

func (c *Circle) Area() float64 {
	return PI * c.radius * c.radius
}

func (r *Rectangle) Type() string {
	return r.name
}

func (c *Circle) Type() string {
	return c.name
}

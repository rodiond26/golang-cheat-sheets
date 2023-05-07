package main

import (
	"fmt"
)

func main() {
	n := make(chan int)
	s := make(chan int)

	// Генерация
	go func() {
		for x := 0; x < 16; x++ {
			n <- x
		}
		close(n)
	}()

	// Возведение в квадрат
	go func() {
		for x := range n {
			s <- x * x
		}
		close(s)
	}()

	// Вывод
	for x := range s {
		fmt.Println(x)
	}
}

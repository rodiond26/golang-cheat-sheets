package main

import (
	"fmt"
	"time"
)

func main() {
	n := make(chan int)
	s := make(chan int)

	// Генерация
	go func() {
		for x := 0; ; x++ {
			n <- x
		}
	}()

	// Возведение в квадрат
	go func() {
		for {
			x, isOpen := <-n
			// Проверка того, что канал не закрыт
			if !isOpen {
				break
			}
			s <- x * x
		}
	}()

	// Вывод
	for {
		time.Sleep(666 * time.Millisecond)
		fmt.Println(<-s)
	}
}

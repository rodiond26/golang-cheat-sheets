package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	timer := time.NewTimer(1 * time.Second)
	quit := make(chan int)
	go jobWithTimeout(timer, quit)

	time.Sleep(2 * time.Second)
	fmt.Println("end")
}

func jobWithTimeout(t *time.Timer, q chan int) {
	time.Sleep(1 * time.Second)
	select {
	case <-t.C:
		fmt.Println("время вышло")
	case <-q:
		if !t.Stop() {
			<-t.C
		}
		fmt.Println("таймер остановлен")
	default:
		fmt.Println("завершение функции")
	}
}

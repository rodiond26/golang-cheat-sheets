package main

import "fmt"

func main() {
	fmt.Println("start main")
	ch := make(chan int)

	go writeChan(ch)

	for i := range ch {
		fmt.Println("chan i :", i)
	}
	fmt.Println("end main")
}

func writeChan(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

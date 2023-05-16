package main

import "fmt"

func foo() {
	panic("паника 100500!")
}

func main() {
	// восстановление после паники
	defer func() {
		err := recover()
		fmt.Println("recovering")
		fmt.Println(err)
	}()
	fmt.Println("start")
	foo()
	fmt.Println("end")
}

package main

import "fmt"

func main() {
	for i := 0; true; i++ {
		fmt.Println("Counter:", i)
		if i > 5 {
			break
		}
	}
}

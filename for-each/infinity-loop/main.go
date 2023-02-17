package main

import "fmt"

func main() {
	c := 0
	for {
		fmt.Println("Counter:", c)
		c++
		if c > 5 {
			break
		}
	}
}

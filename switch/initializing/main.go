package main

import (
	"fmt"
)

func main() {
	for c := 0; c < 20; c++ {
		switch val := c / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non-prime value:", val)
		}
	}
}

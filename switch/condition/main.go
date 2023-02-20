package main

import (
	"fmt"
)

func main() {
	for c := 0; c < 10; c++ {
		switch {
		case c == 0:
			fmt.Println("Zero value")
		case c < 3:
			fmt.Println(c, "is < 3")
		case c >= 3 && c < 7:
			fmt.Println(c, "is >= 3 && < 7")
		default:
			fmt.Println(c, "is >= 7")
		}
	}
}

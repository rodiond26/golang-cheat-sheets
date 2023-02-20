package main

import (
	"fmt"
)

func main() {
	word := "First"
	for i, character := range word {
		switch character {
		case 'F':
			fmt.Println("Uppercase character")
			fallthrough
		case 's':
			fmt.Println("s at position", i)
		case 't':
			fmt.Println("t at position", i)
		}
	}
}

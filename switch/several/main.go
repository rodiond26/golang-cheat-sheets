package main

import (
	"fmt"
)

func main() {
	word := "First"
	for i, character := range word {
		switch character {
		case 'F', 's':
			fmt.Println("F or s at position", i)
		case 't':
			fmt.Println("t at position", i)
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	word := "First"
	for i, character := range word {
		switch character {
		case 'F', 's':
			if character == 's' {
				fmt.Println("Lowercase s at position", i)
				break
			}
			fmt.Println("Uppercase F at position", i)
		case 't':
			fmt.Println("t at position", i)
		default:
			fmt.Println("Character", string(character), "at position", i)
		}
	}
}

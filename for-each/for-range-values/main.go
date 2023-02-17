package main

import "fmt"

func main() {
	word := "first"
	for _, character := range word {
		fmt.Println("Character:", string(character))
	}
}

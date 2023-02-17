package main

import "fmt"

func main() {
	word := "first"
	for i, character := range word {
		fmt.Println("Index:", i, "Character:", string(character))
	}
}

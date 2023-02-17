package main

import "fmt"

func main() {
	word := "first"
	for i := range word {
		fmt.Println("Index:", i)
	}
}

package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	delete(colors, "white")
	fmt.Println(colors)
}

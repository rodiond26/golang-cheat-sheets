package main

import "fmt"

func main() {
	var colors1 map[string]string

	colors2 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors3 := make(map[string]string)

	fmt.Println(colors1)
	fmt.Println(colors2)
	fmt.Println(colors3)
}

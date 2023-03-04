package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printColors(colors)
}

func printColors(colors map[string]string) {
	for color, hexCode := range colors {
		fmt.Printf("Color = %v hex code = %v\n", color, hexCode)
	}
}

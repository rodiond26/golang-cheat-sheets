package main

import "fmt"

func main() {
	var v1 int

	var v2 int = 100

	v3 := 100

	v4 := 100
	v4 = 200

	v5, v6 := 100, 200
	v5, v6 = 200, 100

	var (
		v7 = 100
		v8 = 100
	)

	_ = 100

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
	fmt.Println(v6)
	fmt.Println(v7)
	fmt.Println(v8)
}

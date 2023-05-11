package main

import "fmt"

type Point1 struct {
	X int
	Y int
}

func main() {
	type Point2 struct {
		X int
		Y int
	}

	p1 := Point1{
		X: 5,
		Y: 10,
	}

	p2 := Point2{
		X: 50,
		Y: 100,
	}

	p3 := Point1{1, 2}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	p1.X = 13
	fmt.Println(p1.X)
}

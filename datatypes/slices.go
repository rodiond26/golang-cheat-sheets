package main

import "fmt"

func main() {
	var list []int // []
	l := len(list) // 0
	c := cap(list) // 0

	list = []int{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]
	l = len(list)               // 5
	c = cap(list)               // 5

	list = make([]int, 0, 5) // []
	l = len(list)            // 5
	c = cap(list)            // 5

	// Заполнение среза значениями по умолчанию для данного типа
	list = make([]int, 5, 5) // [0, 0, 0, 0, 0]
	l = len(list)            // 5
	c = cap(list)            // 5

	list = append(list, 1) // [0, 0, 0, 0, 0, 1]
	l = len(list)          // 6
	c = cap(list)          // 10
	fmt.Println(list)
	fmt.Println(l)
	fmt.Println(c)

	list = make([]int, 0, 3) // [] len:0, cap:3
	printSlice(list)
	list = append(list, 1) // [1] len:1, cap:3
	printSlice(list)
	list = append(list, 2) // [1, 2] len:2, cap:3
	printSlice(list)
	list = append(list, 3) // [1, 2, 3] len:3, cap:3
	printSlice(list)
	list = append(list, 4) // [1, 2, 3, 4] len:4, cap:6
	printSlice(list)
}

func printSlice(s []int) {
	fmt.Println("---")
	fmt.Println("s = ", s)
	fmt.Println("len = ", len(s))
	fmt.Println("cap = ", cap(s))
}

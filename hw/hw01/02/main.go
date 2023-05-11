package main

import "fmt"

/*
Решите задачу:
Дан целочисленный массив nums,
выведите на экран true, если любое значение встречается в массиве как минимум дважды,
и false, если каждый элемент различен.
*/

func main() {
	n1 := []int{1, 2, 3, 4, 5}
	r1 := check(n1)
	fmt.Println(r1)
	n2 := []int{1, 2, 3, 4, 3}
	r2 := check(n2)
	fmt.Println(r2)
}

func check(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v] = 0
	}
	return false
}

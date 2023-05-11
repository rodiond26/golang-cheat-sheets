package main

import "fmt"

/*
Решите задачу:
Нам дана строка следующего вида “съешь ещё этих мягких французских булок, да выпей чаю”.
Используя тип данных map посчитайте количество повторений символов в этой строке.
В результате выведите на экран список символ - количество повторений
*/

func main() {
	str1 := "съешь ещё этих мягких французских булок, да выпей чаю"
	m1 := toMap(str1)
	printMap(m1)

	str2 := "аааа ббб вв г дд еее фффф"
	m2 := toMap(str2)
	printMap(m2)
}

func toMap(s string) map[string]int {
	res := make(map[string]int)
	for _, r := range s {
		_, ok := res[string(r)]
		if ok {
			res[string(r)]++
		} else {
			res[string(r)] = 1
		}
	}
	return res
}

func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("[%v] : [%v]\n", k, v)
	}
}

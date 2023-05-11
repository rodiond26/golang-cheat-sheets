package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	var r bool

	r = hasKey(m, "a")
	fmt.Println("r =", r)
	r = hasKey(m, "d")
	fmt.Println("r =", r)
	fmt.Println("end")

	s := "abra cadabra"
	rm := toMap(s)
	sm := toStringIntMap(rm)
	fmt.Println(sm)
}

func hasKey(m map[string]int, k string) bool {
	v, ok := m[k]
	if ok {
		// ключ найден, обрабатываем значение
		fmt.Printf("map has key=%v value=%v\n", k, v)
		return true
	} else {
		// ключ не найден
		fmt.Printf("map doesn't have key=%v\n", k)
		return false
	}
}

func toMap(s string) map[rune]int {
	result := make(map[rune]int)
	// s - строка = массив байт
	for index, value := range s {
		result[value]++
		fmt.Println(index, ":", string(value))
		fmt.Println()
	}
	return result
}

func toStringIntMap(m map[rune]int) map[string]int {
	result := make(map[string]int)
	for key, value := range m {
		result[string(key)] = value
	}
	return result
}

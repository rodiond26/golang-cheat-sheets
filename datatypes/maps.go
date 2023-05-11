package main

import "fmt"

func main() {
	var m1 map[int]bool      // map[]
	var m2 map[string]string // map[]
	m3 := make(map[int]int)  // map[]

	ages := map[string]int{ // map[a:10 b:20 c:30]
		"a": 10,
		"b": 20,
		"c": 30,
	}

	_ = ages["a"]  // 10
	ages["d"] = 40 // map[a:10 a4:40 b:20 c:30 d:40]
	_ = ages["e"]  // 0 - выводит значение по умолчанию для типа элемента, которого нет в отображении
	ages["e"]++    // map[a:10 b:20 c:30 d:40 e:1]

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(ages)
}

package main

import "fmt"

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithError(origin string) (string, error) {
	return "prefix_" + origin, nil
}

func addPrefixWithLength(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(res)
	return
}

func main() {
	myString := addPrefix("my_string")
	fmt.Println(myString) // prefix_my_string

	myString, err := addPrefixWithError("error_string")
	fmt.Println(myString) // prefix_error_string
	fmt.Println(err)      // nil

	n := factorial(20)
	fmt.Println(n)

	var f1, f2 func(s string) int
	f1 = func(s string) int {
		return len(s)
	}
	f2 = func(s string) int {
		return 1
	}

	fmt.Println(f1)
	fmt.Println(f2)

	defer end()
	num := 5
	defer func(x int) {
		fmt.Println(x)
	}(num)
	num = 20
	fmt.Println("end of main function")
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return factorial(n-1) * n
}

// функция adder в своем результате возвращает другую функцию, которая на вход принимает x и возвращает свой результат sum
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func end() {
	fmt.Println("function ends..")
}

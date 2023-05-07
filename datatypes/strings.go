package main

import "fmt"

func main() {
	var myString string
	fmt.Println(myString)

	var hello = "\thello\t"
	fmt.Println(hello)

	var title = `\thello\t`
	fmt.Println(title)

	var b byte = 'c'
	fmt.Println(b)

	var r rune = 'h'
	fmt.Println(r)

	str := "12345"
	fmt.Println(str[0])
	fmt.Println(len(str))

	str2 := "строка"
	fmt.Println(len(str2))

	fmt.Println(str[2:4])
	fmt.Println(str[:4])
	fmt.Println(str[2:])

	isBigger := "aaa" > "aab"
	fmt.Println(isBigger)
}

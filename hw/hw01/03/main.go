package main

import "fmt"

/*
Решите задачу:
Вам дан слайс состоящий из строк.
Необходимо напечатать true тогда и только тогда,
когда все слова в слайсе отсортированы лексикографически относительно друг друга.
*/

func main() {
	s1 := []string{"aa", "ab", "ac", "ad", "ae"}
	r := check(s1)
	fmt.Println(r)
	s2 := []string{"aa", "ab", "ab", "ad", "ae"}
	r2 := check(s2)
	fmt.Println(r2)
	s3 := []string{"aa", "ab", "ad", "ac", "ae"}
	r3 := check(s3)
	fmt.Println(r3)
}

func check(s []string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

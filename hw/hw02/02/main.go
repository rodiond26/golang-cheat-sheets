package main

import "fmt"

//Решите задачу:
//Напишите функцию createList(), на вход которой подается отсортированный по
//возрастанию слайс из целых чисел. Эта функция внутри должна создать цепочку
//структур. По сути это является связным списком.
//
//type ListNode struct {
//	Val int
//	Next *ListNode
//}
//
//Где Val будет принимать значения из переданного в функцию слайса чисел. А в
//Next соответственно должна попасть следующая созданная структура.
//В результате функция должна вернуть структуру ListNode из начала цепочки.
//
//Вторая функция - deleteDuplicates(). На вход она принимает начало вашего
//связного цикла и удаляет из цепочки структуры с одинаковыми Val.
//
//Например, если мы вызовем структуру вот так
//s := createList([]int{1,2,3,3})
//то в s должна вернуться структура ListNode (начала цепочки), у нее значение Val =
//1, а в Next - структура ListNode, у которой значение 2 и т.д все 4 структуры.
//И после вызова deleteDuplicates(s) в который мы передадим наше начало цепочки,
//удаляется последний элемент со значением 3, так как это значение дублируется.
//
//И в результате вернется структура ListNode - начало все той же цепочки, но в ней
//уже на одно значение меньше.

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(list []int) []ListNode {
	res := make([]ListNode, len(list))
	for i := 0; i < len(list); i++ {
		node := ListNode{
			Val: list[i],
		}
		res[i] = node
		if i > 0 {
			res[i-1].Next = &res[i]
		}
	}
	return res
}

func deleteDuplicates(ln []ListNode) []ListNode {
	res := make([]ListNode, 0)
	res = append(res, ln[0])
	for i, j := 1, 0; i < len(ln); i++ {
		if res[j].Val == ln[i].Val {
			continue
		} else {
			res = append(res, ln[i])
			j++
		}
	}
	return res
}

func main() {
	list := []int{1, 2, 3, 3, 4}
	res1 := createList(list)
	res2 := deleteDuplicates(res1)
	fmt.Printf("%v\n", res2)
}

package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contact: contactInfo{
			email:   "alex@mail.com",
			zipCode: 666,
		}}

	alexPointer := &alex
	alexPointer.updateName("Jimmy")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	alexAdd := contactInfo{email: "abc@abc.com", zipCode: 123456}

	alex.contactInfo = alexAdd

	alexPointer := &alex
	alexPointer.updateName("James")
	alex.print()
}

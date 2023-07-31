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

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "anderson",
		contactInfo: contactInfo{
			email:   "alex.anders@gmail.com",
			zipCode: 8080,
		},
	}

	alexPointer := &alex
	alexPointer.updateName("alessi")

	alex.print()
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

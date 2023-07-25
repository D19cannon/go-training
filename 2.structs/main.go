package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "anderson",
		contact: contactInfo{
			email:   "alex.anders@gmail.com",
			zipCode: 8080,
		},
	}

	fmt.Printf("%+v", alex)
}

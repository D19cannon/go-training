package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "alex", lastName: "anderson"}

	fmt.Printf("Person: %+v", alex)

	alex.lastName = "apple"

	fmt.Printf("Person: %+v", alex)
}

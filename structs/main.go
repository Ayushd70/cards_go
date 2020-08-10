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
	ayush := person{
		firstName: "Ayush",
		lastName:  "Dubey",
		contact: contactInfo{
			email:   "ayushdubey70@gmail.com",
			zipCode: 700002,
		},
	}

	fmt.Printf("%+v", ayush)
}

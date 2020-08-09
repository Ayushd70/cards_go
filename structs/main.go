package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	ayush := person{firstName: "Ayush", lastName: "Dubey"}
	fmt.Println(ayush)
}

package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	drinks    bool
}

func main() {
	p1 := person{
		"John",
		"Doe",
		1313,
		true,
	}
	p2 := person{
		"Jane",
		"Doe",
		1001,
		false,
	}

	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

	// accessing fields
	fmt.Printf("%v, %v, %v years. Drinks: %v\n", p1.lastName, p1.firstName, p1.age, p1.drinks)
	fmt.Printf("%v, %v, %v years. Drinks: %v\n", p2.lastName, p2.firstName, p2.age, p2.drinks)
}

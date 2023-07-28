package main

import "fmt"

type person struct {
	firstName string
	age       int
}

func (p person) speak() {
	fmt.Printf("Hello users!! I am %v and my age is %v years\n", p.firstName, p.age)
}

func main() {
	p1 := person{
		firstName: "John",
		age:       42,
	}

	p1.speak()

	(person{
		firstName: "Nichols",
		age:       36,
	}).speak()
}

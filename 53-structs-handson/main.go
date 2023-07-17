package main

import "fmt"

type person struct {
	first              string
	last               string
	favIcecreamFlavors []string
}

func main() {
	p1 := person{
		first:              "John",
		last:               "Doe",
		favIcecreamFlavors: []string{"Chocolate", "Black Current", "Blue berry"},
	}

	p2 := person{
		first:              "Jane",
		last:               "Doe",
		favIcecreamFlavors: []string{"Butter Scotch", "Vanilla", "Coconut"},
	}

	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)
}

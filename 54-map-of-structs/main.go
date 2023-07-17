package main

import "fmt"

type person struct {
	first              string
	last               string
	favIcecreamFlavors []string
}

func main() {
	// personMap := make(map[string]person)
	p1 := person{
		first:              "John",
		last:               "Doe",
		favIcecreamFlavors: []string{"Chocolate", "Black Current", "Blue berry"},
	}

	p2 := person{
		first:              "Crohn",
		last:               "dall-e",
		favIcecreamFlavors: []string{"Butter Scotch", "Vanilla", "Coconut"},
	}

	p3 := person{
		first:              "Gunther",
		last:               "Hart",
		favIcecreamFlavors: []string{"Mango", "Passion-Fruit", "Banana"},
	}

	personMap := map[string]person{
		p1.last: p1,
		p2.last: p2,
		p3.last: p3,
	}

	for k, v := range personMap {
		fmt.Printf("%v\t-\t%#v\n", k, v)
	}
}

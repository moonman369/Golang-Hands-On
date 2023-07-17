package main

import "fmt"

func main() {
	person1 := struct {
		first     string
		friends   []string
		favDrinks []string
	}{
		first:     "John",
		friends:   []string{"Jane", "Bob", "Mark"},
		favDrinks: []string{"Whiskey Sour", "Water", "LIIT"},
	}

	fmt.Printf("%v\n", person1)
}

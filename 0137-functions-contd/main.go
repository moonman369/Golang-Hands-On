package main

import (
	"fmt"
	"strconv"
)

// Demo of Stringer interface
// Any type with methods called `String()` are of type Stringer
// Now when we create specific values of these following types and print them using pln(), they will be printed in the below formats as in their respective String() methods
type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "A Brief History of Time",
	}
	var c count = 42

	fmt.Println(b)
	fmt.Println(c)
}

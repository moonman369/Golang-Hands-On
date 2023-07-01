package main

import "fmt"

func main() {

	/**
	* Go is a statically typed language but datatypes are figured out by the compiler automatically (hence they needn't be always specified)
	*
	 */

	a := 369
	fmt.Printf("a => %d", a)

	b, c, _, e, f := 1, 10, 100, 1000, "Ten Thousand" // Using automatic type assignment where the compiler assigns the types to the vars based on the values being stored in them
	// _ => blank identifier is used to skip values when destructuring from a multiple number of values

	fmt.Printf(`
	b => %d
	c => %d
	e => %d
	f => %s
	`, b, c, e, f)

	// Explicit data type specification during declaration
	var num int = 123131312452
	// var num2 int ....... Using var for zero val initialization
	var str string = "Bich"
	var yn bool = false

	fmt.Printf(`
	num => %d
	str => %s
	yn => %t
	`, num, str, yn)
}

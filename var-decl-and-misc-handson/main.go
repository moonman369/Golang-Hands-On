package main

import "fmt"

// var for zero value
var a int

func main() {
	// Short declaration operator
	b := 42

	// Multiple declaration
	x, y, z := 3.14159, "Hell Yess", true

	// Using var when specificity is reqd
	var c float32 = 2.713

	// Using blank identifier to skip intermediated values
	e, f, _, g := "e", "f", "x", "g"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("x = %v; y = %v; z = %v\n", x, y, z)
	fmt.Println(c)
	fmt.Printf("e = %v; f = %v; g = %v\n\n", e, f, g)

	// Printing variable values and types
	fmt.Printf("b = %v and Type of b = %T\n", b, b)
	fmt.Printf("x = %v and Type of x = %T\n", x, x)
	fmt.Printf("y = %v and Type of y = %T\n", y, y)

}

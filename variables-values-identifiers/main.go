package main

import "fmt"

// VERY IMPORTANT NOTE:
// The Short Declaration Operator i.e `:=` can only be used used inside a function scope

// global := 90  ..........Invalid
var global int = 90

func main() {
	fmt.Printf("global var = %v\n", global)

	/**
	* Go is a statically typed language but datatypes are figured out by the compiler automatically (hence they needn't be always specified)
	*
	 */

	// VERY IMPORTANT NOTE:
	// The Short Declaration Operator i.e `:=` can only be used used inside a function scope

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

	// GO IS A STSTICALLLY TYPED LANGUAGE!!!
	v1 := 42.42
	v2 := "psych"

	fmt.Printf("%v is of type: %T\n", v1, v1)
	fmt.Printf("%v is of type: %T\n", v2, v2)

	// This is an invalid conversion because the types of v1 and v2 as evaluated by the compiler are float64 and string respectively and v2(string) can't be assigned a value of type float64

	// v2 = v1 //Invalid

	// Explicit Type Casting
	var v3 float32 = 42.4242
	v4 := float64(v3)
	fmt.Printf("%v is of type: %T\n", v3, v3)
	fmt.Printf("%v is of type: %T\n", v4, v4)

	// Implicit Type Casting
	var v5 float64 = float64(v3)
	fmt.Printf("%v is of type: %T\n", v5, v5)

}

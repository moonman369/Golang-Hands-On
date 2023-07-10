package main

import "fmt"

func main() {
	// Arrays: Static Aggregate Data Structures. Fixed length

	// Passing length explicitly
	a1 := [5]int{23, 54, 65, 98, 23}
	fmt.Println(a1)

	// Compiler auto detect length
	a2 := [...]string{"Hello", "Gophers", "this", "is", "an", "array"}
	fmt.Println(a2)

	// Only declaring without assigning values
	// Indices are initialized with default value of the given type
	var a3 [5]int
	a3[1], a3[3] = 34, 87
	fmt.Println(a3)

	fmt.Printf("Type of a1: %T\nType of a3: %T\n\n", a1, a3)

	// Arrays can be assigned to each other if they have the same type For example,
	// [n]int can be assigned to [n]int
	// [n]int can't be assigned to [m]int {m != n}
	// [n]int can't be assigned to [n]string
	fmt.Println("a1 before assigning a3", a1)
	a1 = a3
	fmt.Println("a1 after assigning a3", a1)

	fmt.Println("Using in-built len() fn to print the length of a2:", len(a2))

}

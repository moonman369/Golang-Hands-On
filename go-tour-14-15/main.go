package main

import (
	"fmt"
)

func main() {
	// The new variable takes the type inferred from the right hand side

	var x = 77             // Type less decl
	y, z := 42.123, 5+2.3i // Short decl operator

	fmt.Printf("Type of x = %v is %T\n", x, x)
	fmt.Printf("Type of y = %v is %T\n", y, y)
	fmt.Printf("Type of z = %v is %T\n", z, z)

}

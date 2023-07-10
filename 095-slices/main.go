package main

import "fmt"

func main() {
	// No len is required while decl or init
	s1 := []int{12, 34, 65, 34}
	fmt.Println(s1)

	var s2 []int
	fmt.Println(s2)

	// Append to slice using append()
	xi := []int{23, 54, 42, 76, 24, 87}
	fmt.Printf("Before appending: %v\n", xi)
	// Using variadic param to append any arbitrary number of similar type values
	xi = append(xi, 0, 1, 345, 353)
	fmt.Printf("After appending: %v\n", xi)

	// Slicing a slice
	// Works like python slicing x[incl:excl]
	xi1, xi2, xi3 := xi[3:8], xi[3:], xi[:8]

	fmt.Printf("xi[3:8] = %#v\nxi[3:] = %#v\nxi[:8] = %#v\n", xi1, xi2, xi3)
}

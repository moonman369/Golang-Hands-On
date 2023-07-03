package main

import (
	"fmt"
	"math"
	"math/rand"
	"goHandsOn/go-tour-1-3/testpkg2"
)

func main() {
	// fmt.Print()
	// generating random number
	// #1
	fmt.Printf("Random Integer between 0 and 100: %d\n\n", rand.Intn(101))

	// #2
	// Square root using math package
	fmt.Printf("Square root of 1231 = %f\n\n", math.Sqrt(1231))

	// #3
	fmt.Printf("pi =  %f\n", math.Pi)
	fmt.Printf("phi =  %f\n", math.Phi)
	fmt.Printf("e =  %f\n", math.E)
	// fmt.Printf("Random Integer between 0 and 100: %d\n\n", rand.Intn(101))

	something := "Nothing"
	testpkg2.PrintSomething(something)
}

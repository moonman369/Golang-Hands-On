package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe     bool       = true
	MaxInt64 uint64     = 1<<64 - 1
	z        complex128 = cmplx.Sqrt(-5 - 6.9i)
)

func main() {
	fmt.Printf("ToBe ===> %v || NotToBe ===> %v\n", ToBe, !ToBe)
	fmt.Printf("MaxInt64 = %v\n", MaxInt64)
	fmt.Printf("Square Root of z = -5 - 6.9i = %v\n", z)
}

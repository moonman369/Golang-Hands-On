package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 - 6.9i)
)

func main() {
	fmt.Printf("ToBe ===> %v || NotToBe ===> %v\n", ToBe, !ToBe)
	fmt.Printf("MaxInt (64) = %v\n", MaxInt)
	fmt.Printf("Square Root of z = -5 - 6.9i = %v\n", z)
}

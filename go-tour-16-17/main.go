package main

import (
	"fmt"
)

// constants are not assigned any types until their value is required in a context
// numeric constants are high precision values
const (
	Big   = 1 << 100  // Shifting 1 bit to the left by 100 places i.e 1 * 10 ^ 100
	Small = Big >> 99 // Shifting bits of `Big` to right by 99 places i.e 1 << 1 = 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// Bitwise operators
func main() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big), "\n")
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

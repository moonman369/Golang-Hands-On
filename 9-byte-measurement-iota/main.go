package main

import (
	"fmt"
)

// Declaring custom type
type Byte int

const (
	_       = iota
	i1 Byte = 1 << (10 * iota) // 1 << n = 2 ^ n
	i2
	i3
	i4
	i5
	i6
)

func main() {
	fmt.Printf("1 KB = %d \t\t\t\t %bb B\n", i1, i1)
	fmt.Printf("1 MB = %d \t\t\t\t %bb B\n", i2, i2)
	fmt.Printf("1 GB = %d \t\t\t %bb B\n", i3, i3)
	fmt.Printf("1 TB = %d \t\t %bb B\n", i4, i4)
	fmt.Printf("1 PB = %d \t %bb B\n", i5, i5)
	fmt.Printf("1 EB = %d \t %bb B\n", i6, i6)
}

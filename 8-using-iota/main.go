package main

import (
	"fmt"
	"goHandsOn/8-using-iota/testpkg"
)

// iota cannot be used outside a const decl block
const (
	_ = iota
	i1
	i2
	i3
	i4
	i5
	i6
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<i1, 1<<i1)
	fmt.Printf("%d \t %b\n", 1<<i2, 1<<i2)
	fmt.Printf("%d \t %b\n", 1<<i3, 1<<i3)
	fmt.Printf("%d \t %b\n", 1<<i4, 1<<i4)
	fmt.Printf("%d \t %b\n", 1<<i5, 1<<i5)
	fmt.Printf("%d \t %b\n", 1<<i6, 1<<i6)

	fmt.Println(testpkg.Foreign)
	fmt.Println(testpkg.Reverse(testpkg.Foreign))
}

package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210

	fmt.Printf("dec(%v) = %d\n", a, a)
	fmt.Printf("bin(%v) = %b\n", b, b)
	fmt.Printf("dec(%v) = %#x\n", c, c)
}

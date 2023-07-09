package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		var x = rand.Intn(10)
		fmt.Printf("x = %v\n", x)
		switchOver(x)
		fmt.Println()
	}
}

func switchOver(val int) {
	switch val {
	case 0:
		fmt.Println("val is 0")
	case 1:
		fmt.Println("val is 1")
	case 2:
		fmt.Println("val is 2")
	case 3:
		fmt.Println("val is 3")
	case 4:
		fmt.Println("val is 4")
	default:
		fmt.Println("Default case")
	}
}

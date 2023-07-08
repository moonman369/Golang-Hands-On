package main

import (
	"fmt"
	"math/rand"
)

// init is a niladic function i.e it never takes any args
func init() {
	fmt.Println("================INIT-FUNC================")
	printRandRangeSwitch(250)
	fmt.Println("=========================================")
	fmt.Println()
}

func main() {
	for i := 0; i < 10; i++ {
		printRandRangeSwitch(500)
		fmt.Println()
	}
}

func printRandRangeSwitch(limit int) {
	var random int = rand.Intn(limit + 1)
	fmt.Printf("random = %v\n", random)
	switch {
	case random <= 100:
		fmt.Println("Between 0 and 100")
	case random > 100 && random <= 200:
		fmt.Println("Between 101 and 200")
	case random > 200 && random <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("Greater than 250")
	}
}

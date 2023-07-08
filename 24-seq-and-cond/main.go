package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		printRandRange(500)
	}
}

func printRandRange(limit int) {
	var random int = rand.Intn(limit + 1)
	fmt.Printf("random = %v\n", random)
	if random <= 100 {
		fmt.Println("Between 0 and 100")
	} else if random <= 200 {
		fmt.Println("Between 101 and 200")
	} else if random <= 250 {
		fmt.Println("Between 201 and 250")
	} else {
		fmt.Println("Greater than 250")
	}
	fmt.Println()
}

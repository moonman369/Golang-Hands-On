package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count int
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			count++
			fmt.Printf("Iteration %v:\tx is %v!!!\n", i, x)
		}
	}
	fmt.Println("Count is", count)
}

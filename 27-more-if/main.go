package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		checkRand(11)
		fmt.Println()
	}
}

func checkRand(limit int) {
	var x, y int = rand.Intn(limit), rand.Intn(limit)
	fmt.Printf("x = %v; y = %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y both lt 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y both gt 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}
}

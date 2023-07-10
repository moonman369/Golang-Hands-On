package main

import "fmt"

func main() {
	xi := []int{94, 23, 65, 12, 90, 47}
	for idx, val := range xi {
		fmt.Printf("Iterating over a slice.... index = %v; value = %v\n", idx, val)
	}

	m := map[string]int{
		"blud":  0,
		"bills": 2,
		"psych": 193,
		"ligma": 1293801,
		"balls": 98134742,
	}
	for key, val := range m {
		fmt.Printf("Iterating over a map.... key = %v; value = %v\n", key, val)
	}
}

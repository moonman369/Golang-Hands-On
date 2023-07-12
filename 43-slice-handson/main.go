package main

import "fmt"

func main() {
	var xi []int
	for i := 42; i <= 51; i++ {
		xi = append(xi, i)
	}
	fmt.Println(xi)

	for i, v := range xi {
		fmt.Printf("%v \t:\t %v\n", i, v)
	}
}

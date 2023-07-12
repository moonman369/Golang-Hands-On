package main

import "fmt"

func main() {
	var xi []int
	for i := 42; i <= 51; i++ {
		xi = append(xi, i)
	}
	fmt.Println(xi)

	fmt.Println(xi[:5])
	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])
}

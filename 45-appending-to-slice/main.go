package main

import "fmt"

func main() {
	var xi []int
	for i := 42; i <= 51; i++ {
		xi = append(xi, i)
	}

	fmt.Printf("Before appending 52 ---> %#v\n", xi)
	xi = append(xi, 52)
	fmt.Printf("After appending 52 ---> %#v\n", xi)

	xi = append(xi, 53, 54, 55)
	fmt.Printf("After appending 53, 54, 55 ---> %#v\n", xi)

	y := []int{56, 57, 58, 59, 60}
	xi = append(xi, y...)
	fmt.Printf("After appending y ---> %#v\n", xi)

}

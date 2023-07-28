package main

import "fmt"

func main() {
	defer first()
	second()
	third()

	// defer runs in order lifo
	// The last defer func in called first then works its way back to the first defer func
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)
}

func first() {
	fmt.Println("This function has been called first.")
}

func second() {
	fmt.Println("This function has been called second.")
}

func third() {
	fmt.Println("This function has been called third.")
}

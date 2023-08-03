package main

import "fmt"

func main() {
	fmt.Printf("%T\n", Add)
	fmt.Printf("%T\n", Subtract)
	fmt.Printf("%T\n", DoMath)

	x := DoMath(42, 16, Add)
	fmt.Println(x)
	y := DoMath(42, 16, Subtract)
	fmt.Println(y)

}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

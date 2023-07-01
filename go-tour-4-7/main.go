package main

import (
	"fmt"
)

func main() {
	// #4
	sayHello()
	greetAndSayHello("MF")
	fmt.Printf("42 + 24 + 600 = %d\n\n", add(add(42, 24), 600))

	// #5
	s1, s2 := "Ayan", "Mighty"
	s3, s4 := swapStr(s1, s2) //swapStrNamed(s1, s2) ........... gives same result
	fmt.Printf("Swapping strings using multi return func ======> %s, %s =:= %s, %s\n", s1, s2, s3, s4)

	n1, n2 := splitNum(9231980131203)
	fmt.Printf("Splitting 9231980131203 ==> %d + %d", n1, n2)
}

// void return, 0 params
func sayHello() {
	fmt.Println("Hello mf!!")
}

// void return, n params
func greetAndSayHello(name string) {
	fmt.Printf("Hello %s!! Disgusted meeting you!!\n", name)
}

// int return and n params
// if param types are similar, they can be chained to the last param decl for the type
func add(x, y int) int {
	return x + y
}

// multi value return (tuple like return)
func swapStr(s1, s2 string) (string, string) {
	return s2, s1
}

// named return
func splitNum(num int) (x, y int) {
	x = (4 * num) / 9
	y = num - x
	return
}

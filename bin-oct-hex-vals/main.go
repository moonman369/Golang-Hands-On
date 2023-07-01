package main

import "fmt"

func main() {
	a, b, c, d, e, f := 0, 1, 123, 3543, 5345, 38768
	fmt.Printf("%v \t\t %b \t\t %#x \t\t %O\n", a, a, a, a)
	fmt.Printf("%v \t\t %b \t\t %#x \t\t %O\n", b, b, b, b)
	fmt.Printf("%v \t\t %b \t\t %#x \t\t %O\n", c, c, c, c)
	fmt.Printf("%v \t\t %b \t\t %#x \t\t %O\n", d, d, d, d)
	fmt.Printf("%v \t\t %b \t\t %#x \t\t %O\n", e, e, e, e)
	fmt.Printf("%v \t\t %b \t\t %#x \t\t %O\n", f, f, f, f)
}

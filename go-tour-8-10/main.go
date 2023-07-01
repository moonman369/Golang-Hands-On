package main

import (
	"fmt"
)

// Declaring and initializing global vars with zero val
var c, java, python bool
var score1, score2 = 1, 2

func main() {
	fmt.Printf("%v ==> %T\n%v ==> %T\n%v ==> %T\n%v ==> %T\n%v ==> %T\n", c, c, java, java, python, python, score1, score1, score2, score2)
}

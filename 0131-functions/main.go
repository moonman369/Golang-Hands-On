package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// General Syntax:
	// func (r receiverType) identifier (p param/s) (return/s) {code}
	// Receiver is optional. It is used if a func is to be declared as method for a particular type (eg struct)
	// Everything is PASS BY VALUE in Go
	s := aloha("moonman")
	foo(s)
	name, dogAge := getDogYears(7)
	fmt.Printf("%v is %v in dog years\n", name, dogAge)

	s1 := sum(1, 2, 3, 45, 5, 7, 5, 4, 7, 7, 8)
	fmt.Println("The sum is", s1)

	// Unfurling a slice:
	// The above method can also be called by unfurling a suitable type slice in the function args
	x := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s2 := sum(x...)
	fmt.Println("The sum (using slice unfurling) is", s2)

}

func aloha(s string) string {
	return fmt.Sprint("Call me ", s) // Sprint returns a string stream
}

func foo(s string) {
	fmt.Println(s)
}

func getDogYears(ageHuman int) (string, int) {
	names := []string{`zoomie`, `patt`, `tommy`, `barko`, `grim-reaper`, `bully`, `tootie`, `shredder`, `dawglas`, `bernie`}
	return names[rand.Intn(len(names))], ageHuman * 7
}

// Variadic params
// functions in go can be designed to take any arbitrary number (incl 0) of parameters. This can be done using variadic param decl
// IMP: The variadic parameter must be the last incoming parameter in a function and it CANNOT be followed by other params
func sum(numbers /*automatically turns into a slice of given type*/ ...int) int {
	fmt.Printf("%v\n%T\n", numbers, numbers)
	// The param `numbers` can be treated as a slice
	s := 0
	for _, v := range numbers {
		s += v
	}
	return s
}

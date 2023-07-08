// Every executable go prog starts with package main
package main

import (
	"fmt"
	"math/rand"
)

// runs before main()
func init() {
	fmt.Println("This is inside func init()")
}

func main() {
	fmt.Println("This is inside func main()")

	// CONDITIONAL STATEMENTS
	//	|---------IF-ELSE
	//	|---------SWITCH

	// IF-ELSE
	x := 42
	y := 5

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("Greater than or equal to meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("Equal to the meaning of life")
	} else {
		fmt.Println("Greater than meaning of life")
	}

	// Statement-statement idiom
	// The contitional expression can be preceded by a simple statement that is evaluated before the condition check
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("x = %v; z = %v. Hence z >= x\n", x, z)
	} else {
		fmt.Printf("x = %v; z = %v. Hence z < x\n", x, z)
	}
	// There also another thing called comma, ok idiom. Which is used to check if a collection data structure was destructured properly

	// SWITCH
	// Switch supports comparison operators other than `==` like <,>,<=,>=,!=
	// Fallthrough does not occur by default in go -> switch, hence no need for `break`

	switch {
	case x == 42:
		fmt.Println("x is equal to 42")
	case x < 42:
		fmt.Println("x is lesser than 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("Default case for x")
	}

	// Traditional switch-case
	switch x {
	case 40:
		fmt.Println("x is equal to 40")
	case 41:
		fmt.Println("x is equal to 41")
	case 42:
		fmt.Println("x is equal to 42")
	case 43:
		fmt.Println("x is equal to 43")
	default:
		fmt.Println("default case")
	}

	// Fallthrough can be optionally activated using `fallthrough` keyword
	switch x {
	case 40:
		fmt.Println("x is equal to 40. Fallthrough activated.")
		fallthrough
	case 41:
		fmt.Println("x is equal to 41. Fallthrough activated.")
		fallthrough
	case 42:
		fmt.Println("x is equal to 42. Fallthrough activated.")
		fallthrough
	case 43:
		fmt.Println("x is equal to 43. Fallthrough activated.")
		fallthrough
	default:
		fmt.Println("default case")
	}

	// SELECT STATEMENT - USED FOR WRITING CONCURRENT CODE
	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	// Example:
	// select {
	// case v1 := <-ch1:
	// 	fmt.Println("value from channel 1", v1)
	// case v2 := <-ch2:
	// 	fmt.Println("value from channel 2", v2)
	// }

	// LOOPS

	// FOR LOOP
	for i := 0; i < 5; i++ {
		fmt.Printf("Counting to 5.... Current count: %v\n", i)
	}

	// WHILE LOOP
	// The keyword `while` doesn't exist in go. While loops a written using the `for` keyword like the following example
	for y < 12 {
		fmt.Printf("While loop checking.... y = %v.... GTG!\n", y)
		y++
	}

	// We can run for loops without presetting any params like...
	for {
		if y < 25 {
			fmt.Printf("Loop running... y = %v... Good to go!\n", y)
		} else {
			fmt.Println("STOP!!!")
			break
		}
		y++
	}

	// Nested Looping
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}

	// FOR-RANGE used to iterate over a collection datastructures like slices(analog of arrays) and maps(hashmaps)
	xi := []int{42, 283, 234, 746, 89, 345, 67}
	for i, v := range xi {
		fmt.Printf("Iterating over a slice.... index = %v; value = %v\n", i, v)
	}

	m := map[string]int{
		"blud":  0,
		"bills": 2,
		"psych": 193,
		"ligma": 1293801,
		"balls": 98134742,
	}
	for k, v := range m {
		fmt.Printf("Iterating over a map.... key = %v; value = %v\n", k, v)
	}

	s := "whatisareverseexorcismitswhenthedemonasksthepriesttoleavethechildsbody"
	for p, v := range s {
		fmt.Printf("Iterating over a string.... position = %v; value = %v\n", p, rune(v))
	}

}

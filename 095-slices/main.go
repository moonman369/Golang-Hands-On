package main

import "fmt"

func main() {
	// No len is required while decl or init
	s1 := []int{12, 34, 65, 34}
	fmt.Println(s1)

	var s2 []int
	fmt.Println(s2)
}

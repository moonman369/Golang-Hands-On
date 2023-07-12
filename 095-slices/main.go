package main

import (
	"fmt"
	"sort"
)

func main() {
	// No len is required while decl or init
	s1 := []int{12, 34, 65, 34}
	fmt.Println(s1)

	var s2 []int
	fmt.Println(s2)

	// Append to slice using append()
	xi := []int{23, 54, 42, 76, 24, 87}
	fmt.Printf("Before appending: %v\n", xi)
	// Using variadic param to append any arbitrary number of similar type values
	xi = append(xi, 0, 1, 345, 353)
	fmt.Printf("After appending: %v\n", xi)

	// Slicing a slice
	// Works like python slicing x[incl:excl]
	xi1, xi2, xi3 := xi[3:8], xi[3:], xi[:8]

	fmt.Printf("xi[3:8] = %#v\nxi[3:] = %#v\nxi[:8] = %#v\n", xi1, xi2, xi3)

	// Deleting from a slice
	xi4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Using append
	// for deleting item at nth index
	// xi = append(xi[:n], xi[n+1:])
	fmt.Printf("Before deletion of val at idx 4: %#v\n", xi4)
	xi4 = append(xi4[:4], xi4[5:]...) // Destructuring a slice to unpack the elements. Like JS but the ... is added after the slice
	fmt.Printf("After deletion of val at idx 4:  %#v\n", xi4)

	// Creating a slice using `make()` function
	xi5 := make([]int, 0, 10) // (type, init len(size), max length(capacity)) : the max len parameter saves the need to expand the underlying array everytime an overflow element is added to the slice. This saves compile time. If the max len of a slice is known, it is advisable to use make() fn
	fmt.Println(len(xi5))     // returns current length
	fmt.Println(cap(xi5))     // cap returns capacity of the underlying array
	for i := 0; i < 21; i++ {
		xi5 = append(xi5, i)
		fmt.Println("Len =", len(xi5))
		fmt.Println("Cap =", cap(xi5))
	}

	// Multi-dimensional slices
	ay := []string{"Ayan", "Maiti", "Whiskey", "Steak"}
	jd := []string{"John", "Doe", "50% Ethanol-Water Mixture", "Nutrition Sludge"}

	xpeople := [][]string{ay, jd}
	fmt.Println(xpeople)

	// More about underlying arrays in slices
	a := []int{0, 1, 2, 3, 4, 5}
	b := a
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	fmt.Println("=========================")

	a[0] = 99

	// Both a and b changed because both a and b are referencing the same set of values stored once in the memory using pointers
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	// In the above example, b is a shallow copy of a

	// However, to make a deep copy, we must use the copy() function
	c := make([]int, 6)
	copy(c, a)

	a = append(a, 35, 908, 455, 65)

	fmt.Println("a: ", a)
	fmt.Println("c: ", c)
	// c does not change with changes in a. This is because c is pointing to a copy of the underlying array and not the same one as a.

	x := []float64{56, 34, 98, 13324, 4, 0, 99}
	y := make([]float64, 7)
	copy(y, x)

	fmt.Println("x before M1():", x)
	m1 := medianOne(x)
	fmt.Println("x before M1():", x) // Will be sorted in place
	fmt.Println("m1 =", m1)

	fmt.Println("y before M1():", y)
	m2 := medianTwo(y)
	fmt.Println("y before M1():", y) // Will not be sorted in place
	fmt.Println("m2 =", m2)

}

func medianOne(arr []float64) float64 {
	sort.Float64s(arr)
	i := len(arr) / 2
	if len(arr)%2 == 1 {
		return arr[i/2]
	}
	return (arr[i-2] + arr[i]) / 2
}

func medianTwo(arr []float64) float64 {
	c := make([]float64, len(arr))
	copy(c, arr)
	sort.Float64s(c)
	i := len(c) / 2
	if len(c)%2 == 1 {
		return c[i/2]
	}
	return (c[i-2] + c[i]) / 2
}

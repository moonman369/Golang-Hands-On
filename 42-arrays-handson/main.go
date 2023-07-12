package main

import "fmt"

func main() {
	// var arr1 [5]int
	arr := [5]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = (i + 1) * 37
	}

	for i, v := range arr {
		fmt.Printf("%v \t:\t %v\n", i, v)
	}

}

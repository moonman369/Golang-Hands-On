package main

import "fmt"

func main() {
	fmt.Println(sum(1, 45, 7, 9, 9, 8, 6, 5))
}

func sum(nums ...int) (total int) {
	total = 0
	for _, v := range nums {
		total += v
	}
	return
}

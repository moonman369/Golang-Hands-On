package main

import "fmt"

func main() {
	var x int = 50

	for i := 0; i <= x; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}

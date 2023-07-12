package main

import "fmt"

func main() {
	var xi []int
	for i := 42; i <= 51; i++ {
		xi = append(xi, i)
	}
	fmt.Println(xi)

	xi = append(xi[:3], xi[6:]...)
	fmt.Printf("After deleting  xi[3:7] ---> %#v\n", xi)

}

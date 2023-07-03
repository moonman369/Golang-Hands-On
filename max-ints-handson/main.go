package main

import "fmt"

func main() {
	var MaxInt8 int8 = 1<<7 - 1
	var MaxUint8 uint8 = 1<<8 - 1

	fmt.Println(MaxInt8)
	fmt.Println(MaxUint8)
}

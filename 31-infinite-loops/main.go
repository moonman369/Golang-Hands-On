package main

import (
	"fmt"
	"time"
)

// "=========================================================================================="
// "============================CAUTION! THIS IS AN INFINITE LOOP============================="
// "=========================================================================================="

func main() {
	fmt.Println("==========================================================================================")
	fmt.Println("===========================CAUTION! THIS IS AN INFINITE LOOP==============================")
	fmt.Println("==========================================================================================")
	fmt.Println()

	fmt.Println("The infinite loop will start running in 15 seconds.\nIf you want to stop now, then exit the execution with CTRL + C")
	fmt.Println()
	for i := 15; i >= 0; i-- {
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	for i := 0; i < 100; {
		fmt.Print(i)
	}
}

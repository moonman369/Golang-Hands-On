package main

import "fmt"

func main() {
	// xxs := [][]string{}
	// var xxs[][]string
	xxs := make([][]string, 0, 10)
	xxs = append(xxs, []string{"Ayan", "Maiti", "Whiskey", "Steak"}, []string{"John", "Doe", "50% Ethanol-Water Mixture", "Nutrition Sludge"})
	for i, v := range xxs {
		fmt.Printf("%v\t:\t%#v\n", i, v)
	}
}

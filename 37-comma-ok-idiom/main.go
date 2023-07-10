package main

import "fmt"

func main() {
	m := map[string]int{
		"blud":    0,
		"bills":   2,
		"psych":   193,
		"ligma":   1293801,
		"balls":   98134742,
		"grimace": 999999999,
	}

	// comma-ok for assignment statement
	var ok bool
	bludVal, ok := m["gimace"]
	fmt.Println(bludVal, ok)

	// // comma-ok for if statement
	if val, ok := m["gimace"]; !ok {
		fmt.Println("The given key was not found in the map and val returned was m[\"gimace\"] =", val)
	}

	if val, ok := m["ligma"]; ok {
		fmt.Println("The given key was found in the map and val returned was m[\"ligma\"] =", val)
	}
}

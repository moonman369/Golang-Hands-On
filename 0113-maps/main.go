package main

import "fmt"

func main() {
	// varname := map[keyType]valueType {
	// 	k1: v1,
	// 	k2: v2,
	// 	....
	// 	kn: vn
	// }
	m1 := map[string]int{
		"moonman":  22,
		"john-doe": 1231,
		"jane-doe": 2113,
		"kid":      2,
		"dog":      45,
	}
	fmt.Printf("Age of kid is %#v\n", m1["kid"])

	// Creating map using make
	m2 := make(map[string]bool)
	m2["mike"] = true
	m2["timothy"] = false
	m2["jane"] = false
	m2["walter"] = true
	m2["reese"] = true
	fmt.Printf("%#v\n", m2)
	fmt.Println("len(m2) =", len(m2))

	// for-range on a map
	// _, v for just values
	// k, _ for just keys
	for k, v := range m2 {
		fmt.Println(k, "\t:\t", v)
	}

	// Delete element from a map
	m2["delete_this"] = true
	fmt.Printf("%#v\n", m2)
	// Using delete() builtin func
	delete(m2, "delete_this")
	fmt.Printf("%#v\n", m2)

	// Returns the zero value for the respective types of the value field if lookup isn't found. Doesn't panic
	fmt.Printf("Invalid lookup return: %#v\n", m2["khejhsd"])

	m3 := make(map[string]int)
	m3["mike"] = 1093
	m3["timothy"] = 4829737
	m3["jane"] = 13213
	m3["walter"] = 3543
	m3["reese"] = 0

	//comma, ok for maps
	v, ok := m3["jane"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Value was not found after lookup.")
	}

	//comma, ok inline with if
	if v, ok := m3["walter"]; !ok {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Found: ", v)
	}

}

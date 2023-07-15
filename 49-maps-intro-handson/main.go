package main

import "fmt"

func main() {
	m := map[string][]string{}
	m[`bond_james`] = []string{`shaken, not strirred`, `martinis`, `fast cars`}
	m[`moneypenny_jeny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	// fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, ":")
		for i, e := range v {
			fmt.Printf("\t%v : %v\n", i, e)
		}
	}
}

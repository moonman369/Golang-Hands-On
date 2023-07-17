package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	model string
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		model: "Hyundai - Verna",
		color: "Midnight Blue",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		model: "Tata - Nexon",
		color: "Chocolate Brown",
	}

	fmt.Printf("%#v\n", v1)
	fmt.Printf("%#v\n", v2)
}

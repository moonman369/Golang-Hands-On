package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	drinks    bool
}

// Embedded
type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		"John",
		"Doe",
		1313,
		true,
	}
	p2 := person{
		"Jane",
		"Doe",
		1001,
		false,
	}

	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

	// accessing fields
	fmt.Printf("%v, %v, %v years. Drinks: %v\n", p1.lastName, p1.firstName, p1.age, p1.drinks)
	fmt.Printf("%v, %v, %v years. Drinks: %v\n", p2.lastName, p2.firstName, p2.age, p2.drinks)

	// Embedded structs - struct in a struct
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       52,
			drinks:    true,
		},
		ltk: true,
	}
	fmt.Printf("%#v\n", sa1)
	// Automatically decides even if proper tree is not mentioned. Might cause collision if same fieldname is encountered. Fields from the inner structs get promoted to the outher structs
	fmt.Printf("%v, %v, %v years. Drinks: %v. License To Kill: %v\n", sa1.lastName, sa1.firstName, sa1.age, sa1.drinks, sa1.ltk)
	fmt.Printf("%v, %v, %v years. Drinks: %v. License To Kill: %v\n", sa1.person.lastName, sa1.person.firstName, sa1.person.age, sa1.person.drinks, sa1.ltk)

	// Anonymous  structs
	// We can skip the struct declaration and straight up create a struct type variable while declaring the struct type in the same statement. This is great for a one time use struct and it required no name for that particular type
	// Normal struct decl: s := structType {
	// 	field1: val1,
	// 	.....
	// }

	// Anonymous struct decl:
	// s := struct {
	// 	field1 type1
	// 	....
	// } {
	// 	field1: value1,
	// 	...
	// }

	car1 := struct {
		company  string
		model    string
		topSpeed int
	}{
		company:  "Hyundai",
		model:    "Verna",
		topSpeed: 180, // If a field is skipped, the default/zero value is assigned to field
	}
	fmt.Printf("%#v\n", car1)

}

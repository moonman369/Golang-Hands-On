package main

import (
	"fmt"

	puppy "github.com/moonman369/Puppy-Go-Dependency"
)

func main() {
	var n int = 9
	singleBark, nBarks, singleGrown, nGrown := puppy.Bark(), puppy.Barks(n), puppy.BigBark(), puppy.BigBarks(n)

	fmt.Println("Puppy name before setting: ", puppy.PuppyName)
	puppy.SetPuppyName("Cosmic Annihilation")
	fmt.Println("Puppy name after setting: ", puppy.PuppyName)
	fmt.Println()

	fmt.Println("Puppy age before setting: ", puppy.PuppyAge)
	puppy.SetPuppyAge(int8(n))
	fmt.Println("Puppy age after setting: ", puppy.PuppyAge)
	fmt.Println()

	fmt.Println("Puppy favorite snack before setting: ", puppy.FavSnack)
	puppy.SetFavSnack("5 Year olds")
	fmt.Println("Puppy favorite snack after setting: ", puppy.FavSnack)
	fmt.Println()

	fmt.Printf("`puppy.Bark()` call returned ==> %v\n`puppy.Barks(n int)` call returned ==> %v\n`puppy.BigBark()` call returned ==> %v\n`puppy.BigBarks(n int)` call returned ==> %v\n", singleBark, nBarks, singleGrown, nGrown)

}

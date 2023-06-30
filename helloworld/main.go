package main

import "fmt"

func main() {
	// Rawstring litarals. Preserves tabs, newlines, etc.
	fmt.Println(`
	UTF-8 saves space.
	Int UTF-8, common characters like "C" take 8 bits, 
	while rare characters like "🎎🎍🎁" take 32 bits.
	Other characters take 16 or 24 bits.
	A blog post like this one takes about
	four times less space in UTF-8 than it would in UTF-32.
	As a result, it loads faster.
	`)
}

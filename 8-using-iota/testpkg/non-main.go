package testpkg

// package scope testing
var Foreign int = 369

func Reverse(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + (x % 10)
		x /= 10
	}
	return r
}

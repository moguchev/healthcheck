// pow.go
package main

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y int) int {
	if y == 0 {
		return 1
	} else if y < 0 {
		return 0
	}

	var res = 1
	for i := 0; i < y; i++ {
		res *= x
	}

	return res
}

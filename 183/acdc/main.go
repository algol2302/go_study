// Package acdc born to be wild
package acdc

// Sum returns sum of imput elements
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

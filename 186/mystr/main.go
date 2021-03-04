package mystr

import "strings"

// Cat is concat function
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// Join is like python join function
func Join(xs []string) string {
	return strings.Join(xs, " ")
}

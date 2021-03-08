package word

import "strings"

// UseCount returns a map with counter of unique words for string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns a total count of words in string
func Count(s string) int {
	xs := strings.Fields(s)
	var total int
	for range xs {
		total++
	}
	return total
}

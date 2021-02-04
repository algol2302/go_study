package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 5, 6, 8, 42, 2, 56, 77, 1}
	xs := []string{"James Bond", "M", "Q", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}

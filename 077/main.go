package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i, xs := range x {
		fmt.Println("record: ", i)
		for j, v := range xs {
			fmt.Printf("\t index: %v \t value: %v\n", j, v)
		}
	}
}

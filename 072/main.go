package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, v := range s {
		fmt.Println(index, v)
	}
	fmt.Printf("%T", s)
}

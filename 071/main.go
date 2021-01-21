package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for index, v := range a {
		fmt.Println(index, v)
	}
	fmt.Printf("%T", a)
}

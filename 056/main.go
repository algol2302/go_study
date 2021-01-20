package main

import "fmt"

func main() {
	condition := "c"
	if condition == "a" {
		fmt.Println(condition)
	} else if condition == "b" {
		fmt.Println(condition)
	} else {
		fmt.Println("something else")
	}
}

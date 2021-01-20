package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("true - always printed")
	case false:
		fmt.Println("false - never printed")
	}

}

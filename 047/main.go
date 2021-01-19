package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true 1")
		fallthrough
	case (2 == 3):
		fmt.Println("not true 2")
		fallthrough
	case (6 == 6):
		fmt.Println("prints")
		fallthrough
	default:
		fmt.Println("default case")
	}

	switch "A" {
	case "Moneypenny":
		fmt.Println("Moneypenny")
	case "Bond":
		fmt.Println("Bond")
	default:
		fmt.Println("default case")
	}

	n := "A"
	switch n {
	case "A", "Moneypenny", "Q":
		fmt.Println("A, Moneypenny or Q")
	case "Bond":
		fmt.Println("Bond")
	default:
		fmt.Println("default case")
	}

}

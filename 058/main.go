package main

import "fmt"

func main() {
	favSport := "asd"

	switch favSport {
	case "karate":
		fmt.Println("karate")
	case "soccer":
		fmt.Println("soccer")
	default:
		fmt.Println("choose sport")
	}
}

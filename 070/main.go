package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneymenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Q"])

	delete(m, "James")
	fmt.Println(m)

	delete(m, "A")
	fmt.Println(m)
}

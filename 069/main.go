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

	v, ok := m["Q"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Q"] = 33

	if v, ok := m["Q"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}

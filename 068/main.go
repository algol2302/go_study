package main

import "fmt"

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

	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}

}

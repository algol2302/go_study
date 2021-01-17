package main

import "fmt"

func main() {
	// 	==
	// <=
	// >=
	// !=
	// <
	// >

	g := (42 == 43)
	h := (1 <= 2)
	i := (2 >= 3)
	j := (2 != 3)
	k := (2 < 3)
	l := (3 > 4)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

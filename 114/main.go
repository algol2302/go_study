package main

import "fmt"

func main() {
	x := 0
	fmt.Println("before", &x)
	fmt.Println("before", x)
	foo(&x)
	fmt.Println("after", &x)
	fmt.Println("after", x)
}

func foo(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}

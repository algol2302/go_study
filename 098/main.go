package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	x := bar()
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

func foo() string {
	return "Hello world"
}

func bar() func() int {
	return func() int { return 451 }
}

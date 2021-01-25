package main

import "fmt"

func main() {
	foo()
	bar("James")

	s1 := woo("Moneypenny")
	fmt.Println(s1)

	x, y := mouse("Yan", "Fleming")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("hello from foo!")
}

func bar(s string) {
	fmt.Println("Hello,", s, "!")
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s, "!")
}

func mouse(fn string, ln string) (string, bool) {
	return fmt.Sprint(fn, " ", ln, `, says "Hello!"`), true
}

package main

import "fmt"

func main() {
	f1 := func() {
		fmt.Println("my first func expression")
	}
	f1()

	f2 := func(x int) {
		fmt.Println("my second func expression", x)
	}
	f2(1984)
}

package main

import (
	"fmt"
)

type customErr struct {
	error string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("there is a custom error: %v", ce.error)
}

func main() {
	c1 := customErr{
		error: "Ahtung!",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}

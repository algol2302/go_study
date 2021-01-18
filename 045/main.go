package main

import (
	"fmt"
)

func main() {
	x := 41
	if x == 40 {
		fmt.Println("our value is 40")
	} else if x == 42 {
		fmt.Println("our value is 42")
	} else {
		fmt.Println("our value is not 40 or 42")
	}
}

package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm an anomymous function")
	}()
}

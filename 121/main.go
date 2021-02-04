package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	fmt.Fprintln(os.Stdout, "Hello!")
	io.WriteString(os.Stdout, "Hellooo!")
}

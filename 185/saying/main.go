package saying

import "fmt"

// Greet is help simple function
func Greet(s string) string {
	return fmt.Sprintf("Welcome my dear %v", s)
}

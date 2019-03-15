package helloworld

import "fmt"

// Hello returns a friendly greeting
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

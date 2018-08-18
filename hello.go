// Package go_demo_lib contains utility functions for Go-Demo-3 application.
package go_demo_lib

import "fmt"

// Hello returns it's argument in the form of a greeting
// For example, supplying Joost, will return "Hello Joost".
func Hello(nameToGreet string) string {
	return fmt.Sprintf("Hello %s", nameToGreet)
}
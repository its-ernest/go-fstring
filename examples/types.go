package main

import (
	"fmt"
	"github.com/its-ernest/go-fstring"
)

func main() {
	f := fstring.New()

	vars := map[string]any{
		"name":     "Ama",
		"age":      25,
		"active":   true,
		"balance":  99.75,
		"metadata": struct{ ID int }{ID: 123},
	}

	msg := f.Str("{name} is {age} years old, active={active}, balance={balance}, metadata={metadata}", vars)
	fmt.Println(msg)
	// Output:
	// Ama is 25 years old, active=true, balance=99.75, metadata={123}
}
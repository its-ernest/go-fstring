package main

import (
	"fmt"
	"github.com/its-ernest/go-fstring"
)

func main() {
	f := fstring.New()

	vars := map[string]any{
		"name":   "Kofi",
		"action": "coding",
	}

	msg := f.Str("{name} is currently {action} in Go!", vars)
	fmt.Println(msg)
	// Output: Kofi is currently coding in Go!
}
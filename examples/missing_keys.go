package main

import (
	"fmt"
	"github.com/its-ernest/go-fstring"
)

func main() {
	f := fstring.New()

	vars := map[string]any{
		"name": "Kojo",
	}

	msg := f.Str("Hello {name}, your task is {task}", vars)
	fmt.Println(msg)
	// Output: Hello Kojo, your task is {task}
}
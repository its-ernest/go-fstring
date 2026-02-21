package main

import (
	"fmt"
	"github.com/its-ernest/go-fstring"
)

func main() {
	f := fstring.New()

	// Dynamic map updates
	vars := make(map[string]any)

	for i := 1; i <= 3; i++ {
		vars["user"] = fmt.Sprintf("User%d", i)
		vars["score"] = i * 10

		msg := f.Str("{user} has a score of {score}", vars)
		fmt.Println(msg)
	}
	// Output:
	// User1 has a score of 10
	// User2 has a score of 20
	// User3 has a score of 30
}
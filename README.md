# go-fstring

Python-style f-string interpolation for Go using maps.  
Build dynamic strings with `{VarName}` placeholders, minimal boilerplate.

---

## Features

- Map-based placeholders: `{VarName}` replaced with map values
- Works with any type (`string`, `int`, `float`, `bool`, structs via `fmt.Sprint`)
- Stateless and thread-safe
- Leaves unknown placeholders intact
- Simple,  less boilerplate API

---

## Installation

```bash
go get github.com/its-ernest/go-fstring
```

## Example

```go
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
		"age":    30,
	}

	msg := f.Str("{name} is currently {action} in Go and is {age} years old.", vars)
	fmt.Println(msg)
	// Output: Kofi is currently coding in Go and is 30 years old.
}
```

**Missing variables:**

```go
msg2 := f.Str("{name} is {action} doing {missing}.", vars)
fmt.Println(msg2)
// Output: Kofi is coding doing {missing}.
```

## Contributing
*go-fstring* is open to contributions. To contribute, start this repo and follow the [CONTRIBUTING.md](CONTRIBUTING.md)
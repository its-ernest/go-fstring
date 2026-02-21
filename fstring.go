// Package fstring provides Python-style f-string interpolation using maps.
package fstring

import (
	"bytes"
	"fmt"
	"regexp"
)

// F is the f-string context. Stateless in v1.
type F struct{}

// New returns a new F instance.
func New() *F {
	return &F{}
}

// Str formats the template using the provided variables map.
// Placeholders should be in the form {VarName}.
func (f *F) Str(template string, vars map[string]any) string {
	return format(template, vars)
}

// format does the actual replacement
func format(template string, vars map[string]any) string {
	re := regexp.MustCompile(`\{(\w+)\}`)
	var buf bytes.Buffer

	lastIndex := 0
	matches := re.FindAllStringSubmatchIndex(template, -1)

	for _, m := range matches {
		// Append text before placeholder
		buf.WriteString(template[lastIndex:m[0]])

		// Extract key
		key := template[m[2]:m[3]]

		// Replace with value if present
		if val, ok := vars[key]; ok {
			buf.WriteString(fmt.Sprint(val))
		} else {
			buf.WriteString("{" + key + "}") // leave as-is if missing
		}

		lastIndex = m[1]
	}

	// Append remaining text
	buf.WriteString(template[lastIndex:])
	return buf.String()
}
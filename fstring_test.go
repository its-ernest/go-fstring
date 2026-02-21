package fstring

import "testing"

func TestStrBasic(t *testing.T) {
	f := New()

	vars := map[string]any{
		"name":   "Kofi",
		"action": "coding",
		"age":    30,
	}

	expected := "Kofi is coding in Go and is 30 years old."
	result := f.Str("{name} is {action} in Go and is {age} years old.", vars)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestMissingVar(t *testing.T) {
	f := New()

	vars := map[string]any{
		"name": "Kofi",
	}

	expected := "Kofi is {action}"
	result := f.Str("{name} is {action}", vars)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestVariousTypes(t *testing.T) {
	f := New()

	vars := map[string]any{
		"str":   "hello",
		"int":   42,
		"float": 3.14,
		"bool":  true,
	}

	expected := "hello 42 3.14 true"
	result := f.Str("{str} {int} {float} {bool}", vars)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
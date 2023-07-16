package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := map[string]bool{
		"(ba":     false,
		"(ba])":   false,
		"(foo)":   true,
		"()[]{}":  true,
		"[{foo}]": true,
		"[{()}]":  true,
	}

	for s, expected := range tests {
		result := isValid(s)

		if result != expected {
			t.Logf("input: %s", s)
			t.Fatalf("unexpected value: expecting %v, got %v", expected, result)
		}
	}
}

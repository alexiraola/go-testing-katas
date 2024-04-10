package main

import "testing"

func TestEmptyStringShouldReturnEmptyString(t *testing.T) {
	result := camelcase("")

	if result != "" {
		t.Fatalf(`expected %q, got %q`, "", result)
	}
}

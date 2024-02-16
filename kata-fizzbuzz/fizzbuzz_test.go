package main

import "testing"

func TestOneShouldReturnOne(t *testing.T) {
	result := fizzbuzz(1)

	if result != "1" {
		t.Fatalf(`result should be "1", got %q`, result)
	}
}

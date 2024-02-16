package main

import "testing"

func TestOneShouldReturnOne(t *testing.T) {
	result := fizzbuzz(1)

	if result != "1" {
		t.Fatalf(`result should be "1", got %q`, result)
	}
}

func TestThreeShouldReturnFizz(t *testing.T) {
	result := fizzbuzz(3)

	if result != "fizz" {
		t.Fatalf(`result should be "fizz", got %q`, result)
	}
}

func TestFiveShouldReturnBuzz(t *testing.T) {
	result := fizzbuzz(5)

	if result != "buzz" {
		t.Fatalf(`result should be "buzz", got %q`, result)
	}
}

func TestFifteenShouldReturnFizzbuzz(t *testing.T) {
	result := fizzbuzz(15)

	if result != "fizzbuzz" {
		t.Fatalf(`result should be "fizzbuzz", got %q`, result)
	}
}

func TestDivisibleByThreeShouldReturnFizz(t *testing.T) {
	result := fizzbuzz(9)

	if result != "fizz" {
		t.Fatalf(`result should be "fizz", got %q`, result)
	}
}

func TestDivisibleByFiveShouldReturnBuzz(t *testing.T) {
	result := fizzbuzz(10)

	if result != "buzz" {
		t.Fatalf(`result should be "buzz", got %q`, result)
	}
}

func TestDivisibleByThreeAndFiveShouldReturnFizzbuzz(t *testing.T) {
	result := fizzbuzz(30)

	if result != "fizzbuzz" {
		t.Fatalf(`result should be "fizzbuzz", got %q`, result)
	}
}

func TestNotDivisibleByThreeOrFiveShouldReturnTheSameNumber(t *testing.T) {
	result := fizzbuzz(23)

	if result != "23" {
		t.Fatalf(`result should be "23", got %q`, result)
	}
}

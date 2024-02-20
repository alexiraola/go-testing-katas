package main

import "testing"

func TestEmptyStringShouldReturnEmptyString(t *testing.T) {
	result := camelcase("")

	if result != "" {
		t.Fatalf(`expected %q, got %q`, "", result)
	}
}

func TestStringStartingWithUppercaseShouldReturnSameString(t *testing.T) {
	result := camelcase("Foo")

	if result != "Foo" {
		t.Fatalf(`expected %q, got %q`, "Foo", result)
	}
}

func TestStringStartingWithLowercaseShouldReturnSameStringUppercased(t *testing.T) {
	result := camelcase("foo")

	if result != "Foo" {
		t.Fatalf(`expected %q, got %q`, "Foo", result)
	}
}

func TestWordsSeparatedByWhitespaceShouldJoin(t *testing.T) {
	result := camelcase("Foo Bar")

	if result != "FooBar" {
		t.Fatalf(`expected %q, got %q`, "FooBar", result)
	}
}

func TestWordsSeparatedByHyphensShouldJoin(t *testing.T) {
	result := camelcase("Foo_Bar-Foo")

	if result != "FooBarFoo" {
		t.Fatalf(`expected %q, got %q`, "FooBar", result)
	}
}

func TestLowercaseWordsSeparatedByHyphensShouldJoinAndUppercase(t *testing.T) {
	result := camelcase("foo_bar foo-bar")

	if result != "FooBarFooBar" {
		t.Fatalf(`expected %q, got %q`, "FooBar", result)
	}
}

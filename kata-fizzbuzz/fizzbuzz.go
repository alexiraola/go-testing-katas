package main

import "fmt"

func fizzbuzz(number int) string {
	isDivisibleBy := func(divisor int) bool {
		return number%divisor == 0
	}

	if isDivisibleBy(15) {
		return "fizzbuzz"
	}
	if isDivisibleBy(3) {
		return "fizz"
	}
	if isDivisibleBy(5) {
		return "buzz"
	}
	return fmt.Sprint(number)
}

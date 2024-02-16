package library

func sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func avg(numbers []int) float64 {
	return float64(sum(numbers)) / float64(len(numbers))
}

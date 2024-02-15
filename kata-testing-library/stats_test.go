package library

import "testing"

func TestSum(t *testing.T) {
	sum := sum([]int{1, 2, 3})

	if sum != 6 {
		t.Fatalf(`sum = %d`, sum)
	}
}

func TestAverage(t *testing.T) {
	avg := avg([]int{1, 2, 3})

	if avg != 2 {
		t.Fatalf(`avg = %f`, avg)
	}
}

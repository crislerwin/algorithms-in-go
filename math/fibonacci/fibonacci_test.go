package fibonacci

import (
	"testing"

	"www.github.com/crislerwin/algorithms-in-go/dynamic"
)

func getTests() []struct {
	name string
	n    uint
	want uint
} {
	tests := []struct {
		name string
		n    uint
		want uint
	}{
		{"Fibonacci of 0-th number == 0", 0, 0},
		{"Fibonacci of 1-st number == 1", 1, 1},
		{"Fibonacci of 2-nd number == 1", 2, 1},
		{"Fibonacci of 3-rd number == 2", 3, 2},
		{"Fibonacci of 4-th number == 3", 4, 3},
		{"Fibonacci of 5-th number == 5", 5, 5},
		{"Fibonacci of 6-th number == 8", 6, 8},
		{"Fibonacci of 7-th number == 13", 7, 13},
		{"Fibonacci of 8-th number == 21", 8, 21},
		{"Fibonacci of 9-th number == 34", 9, 34},
		{"Fibonacci of 10-th number == 55", 10, 55},
		{"Fibonacci of 20-th number == 6765", 20, 6765},
	}
	return tests
}

func TestMatrix(t *testing.T) {
	tests := getTests()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := dynamic.NthFibonnacci(test.n); got != test.want {
				t.Errorf("Return value - %v, want %v", got, test.want)
			}

		})
	}
}

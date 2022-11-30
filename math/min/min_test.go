package min

import "testing"

func getTestCases() []struct {
	name    string
	base    int
	numbers []int
	min     int
} {
	var tests = []struct {
		name    string
		base    int
		numbers []int
		min     int
	}{
		{"Minimun of [128, 127], min = 117", 8, []int{128, 127}, 127},
		{"Minimun of [5], min = 5", 32, []int{5}, 5},
		{"Minimum of [-8, 32, 64, -1, 0], min = -8", 64, []int{-8, 32, 64, -1, 0}, -8},
		{"Minimum of [1, 2, 3, 4, 5], min = 1", 32, []int{1, 2, 3, 4, 5}, 1},
		{"Minimum of [1024, 512, 256, 333, 777], min = 256", 64, []int{1024, 512, 256, 333, 777}, 256},
		{"Minimum of [-9223372036854770001, -9223372036854770000, 256, 333, 777], min = 256", 64, []int{-9223372036854770001, -9223372036854770000, 256, 333, 777}, -9223372036854770001},
	}
	return tests

}

func TestMin(t *testing.T) {
	for _, test := range getTestCases() {
		t.Run(test.name, func(t *testing.T) {
			min := Int(test.numbers...)
			if min != test.min {
				t.Errorf("Expected %d, got %d", test.min, min)
			}
		})
	}
}

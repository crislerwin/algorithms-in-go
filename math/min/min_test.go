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

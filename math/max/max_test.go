package max

import "testing"

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		left  int
		rigth int
		max   int
	}{
		{name: "Left is max", left: 10, rigth: 5, max: 10},
		{name: "Right is max", left: 5, rigth: 10, max: 10},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			returnedMax := Int(test.left, test.rigth)
			if returnedMax != test.max {
				t.Errorf("Failed test %s. Expected %s\n\tleft: %v, right: %v, max: %v but received: %v", test.name, test.name, test.left, test.rigth, test.max, returnedMax)
			}
		})

	}
}

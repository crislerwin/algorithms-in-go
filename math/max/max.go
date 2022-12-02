package max

import "www.github.com/crislerwin/algorithms-in-go/constraints"

func Int[T constraints.Integer](values ...T) T {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}

	}
	return max
}

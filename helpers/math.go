package helpers

import "math"

// Min finds the smallest of the values
func Min(n ...int) (min int) {
	min = math.MaxInt64
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	return
}

// Max finds the largest of the values
func Max(n ...int) (max int) {
	max = math.MinInt64
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return
}

package main

import "math"

// PartTwo sovles the second part of the problem.
func PartTwo(input Input, duration int) (solution Solution) {
	m := make(map[Reindeer]int)
	for _, r := range input {
		m[r] = 0
	}

	for i := 1; i <= duration; i++ {
		leadDistance := -1
		for r := range m {
			d := r.DistanceAfter(i)
			if d > leadDistance {
				leadDistance = d
			}
		}
		for r := range m {
			if r.DistanceAfter(i) == leadDistance {
				m[r]++
			}
		}
	}

	max := math.MinInt64
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return Solution(max)
}

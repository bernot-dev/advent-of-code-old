package main

import (
	"math"

	"github.com/bernot-dev/advent-of-code/helpers"
)

// DistanceAfter .
func (r *Reindeer) DistanceAfter(s int) (d int) {
	combined := r.Duration + r.Rest

	full := s / combined
	d = full * r.Speed * r.Duration

	partial := s % combined
	additional := helpers.Min(partial, r.Duration)
	d += additional * r.Speed

	return d
}

// PartOne solves the first part of the problem.
func PartOne(input Input, duration int) (solution Solution) {
	max := math.MinInt64
	for _, r := range input {
		d := r.DistanceAfter(duration)
		if d > max {
			max = d
		}
	}
	return Solution(max)
}

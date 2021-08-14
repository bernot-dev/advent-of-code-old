package main

import (
	"math"
)

// Weight .
func (g Graph) Weight(a, b string) int {
	return g[Edge{a, b}] + g[Edge{b, a}]
}

// Heaviest .
func (g Graph) Heaviest(first, current string, remaining []string) int {
	if len(remaining) == 0 {
		return g.Weight(current, first)
	}

	max := math.MinInt64
	for i := range remaining {
		remainingCopy := make([]string, 0, len(remaining)-1)
		remainingCopy = append(remainingCopy, remaining[:i]...)
		if i+1 < len(remaining) {
			remainingCopy = append(remainingCopy, remaining[i+1:]...)
		}

		curWeight := g.Weight(current, remaining[i])
		nextWeight := curWeight + g.Heaviest(first, remaining[i], remainingCopy)
		if nextWeight > max {
			max = nextWeight
		}
	}
	return max
}

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	remaining := input.Names[1:]
	return Solution(input.Graph.Heaviest(input.First, input.First, remaining))
}

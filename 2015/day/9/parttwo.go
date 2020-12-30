package main

import (
	"github.com/bernot-dev/advent-of-code/helpers"
)

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	cities := make([]string, 0, len(input))
	for c := range input {
		cities = append(cities, c)
	}
	return Solution(input.Distance(helpers.Max, cities))
}

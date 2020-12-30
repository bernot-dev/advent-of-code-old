package main

import (
	"github.com/bernot-dev/advent-of-code/helpers"
)

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	cities := make([]string, 0, len(input))
	for c := range input {
		cities = append(cities, c)
	}
	return Solution(input.Distance(helpers.Min, cities))
}

// Distance returns the distance to visit all points
func (in *Input) Distance(fn func(...int) int, cities []string) (result int) {
	ds := make([]int, len(cities))
	for i, cs := range Permutations(cities) {
		ds[i] = in.remainingDistance(fn, cs[0], cs[1:])
	}
	return fn(ds...)
}

// remainingDistance returns the distance to visit the remaining points
func (in *Input) remainingDistance(fn func(...int) int, city string, remaining []string) (result int) {
	if len(remaining) == 1 {
		return (*in)[city][remaining[0]]
	}
	ds := make([]int, len(remaining))
	for i, cs := range Permutations(remaining) {
		ds[i] = (*in)[city][cs[0]] + in.remainingDistance(fn, cs[0], cs[1:])
	}
	return fn(ds...)
}

// Permutations .
func Permutations(s []string) [][]string {
	ps := make([][]string, len(s))
	for i := range ps {
		ps[i] = make([]string, len(s))
		copy(ps[i], s)
		ps[i][0], ps[i][i] = ps[i][i], ps[i][0]
	}
	return ps
}

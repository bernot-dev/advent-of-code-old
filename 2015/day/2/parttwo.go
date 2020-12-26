package main

import (
	"sort"
)

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	return
}

// Ribbon determines teh amount of ribbon needed for the box
func (p *Present) Ribbon() int {
	dimensions := []int{p.h, p.l, p.w}
	sort.Sort(sort.IntSlice(dimensions))
	ribbon := 2 * (dimensions[0] + dimensions[1])
	bow := p.Volume()
	return ribbon + bow
}

// Volume calculates the volume of the box
func (p *Present) Volume() int {
	return p.h * p.l * p.w
}

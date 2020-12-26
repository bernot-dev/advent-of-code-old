package main

import "github.com/bernot-dev/advent-of-code/helpers"

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	return
}

// Wrapping determines how much wrapping paper is required for the present
func (p *Present) Wrapping() int {
	s1 := p.l * p.w
	s2 := p.w * p.h
	s3 := p.h * p.l
	extra := helpers.Min(s1, s2, s3)
	return 2*(s1+s2+s3) + extra
}

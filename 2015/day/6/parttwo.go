package main

import "github.com/bernot-dev/advent-of-code/helpers"

// DimGrid keeps track of brightness
type DimGrid map[Coordinate]int

// PartTwo solves the first part of the problem.
func PartTwo(input Input) (solution Solution) {
	g := make(DimGrid)
	for _, ins := range input {
		g.FollowInstruction(ins)
	}
	for _, v := range g {
		solution += Solution(v)
	}
	return
}

// FollowInstruction provided
func (g *DimGrid) FollowInstruction(ins Instruction) {
	var action func(x, y int)
	switch ins.command {
	case "turn on":
		action = g.Brighten
	case "turn off":
		action = g.Dim
	case "toggle":
		action = g.Brighten2
	}

	for i := ins.start.x; i <= ins.end.x; i++ {
		for j := ins.start.y; j <= ins.end.y; j++ {
			action(i, j)
		}

	}
}

// Dim light
func (g *DimGrid) Dim(x, y int) {
	val := (*g)[Coordinate{x, y}]
	(*g)[Coordinate{x, y}] = helpers.Max(val-1, 0)
}

// Brighten light
func (g *DimGrid) Brighten(x, y int) {
	(*g)[Coordinate{x, y}]++
}

// Brighten2 light
func (g *DimGrid) Brighten2(x, y int) {
	(*g)[Coordinate{x, y}] += 2
}

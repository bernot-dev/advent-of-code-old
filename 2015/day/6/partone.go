package main

// Grid of lights
type Grid map[Coordinate]bool

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	g := make(Grid)
	for _, ins := range input {
		g.FollowInstruction(ins)
	}
	return Solution(len(g))
}

// FollowInstruction provided
func (g *Grid) FollowInstruction(ins Instruction) {
	var action func(x, y int)
	switch ins.command {
	case "turn on":
		action = g.TurnOn
	case "turn off":
		action = g.TurnOff
	case "toggle":
		action = g.Toggle
	}

	for i := ins.start.x; i <= ins.end.x; i++ {
		for j := ins.start.y; j <= ins.end.y; j++ {
			action(i, j)
		}

	}
}

// TurnOff light
func (g *Grid) TurnOff(x, y int) {
	delete(*g, Coordinate{x, y})
}

// TurnOn light
func (g *Grid) TurnOn(x, y int) {
	(*g)[Coordinate{x, y}] = true
}

// Toggle light
func (g *Grid) Toggle(x, y int) {
	if (*g)[Coordinate{x, y}] {
		delete(*g, Coordinate{x, y})
	} else {
		(*g)[Coordinate{x, y}] = true
	}
}

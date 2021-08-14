package main

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	remaining := append(input.Names[1:], "Adam")
	return Solution(input.Graph.Heaviest(input.First, input.First, remaining))
}

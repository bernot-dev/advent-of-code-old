package main

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	for i := 0; i < 50; i++ {
		input = LookAndSay(input)
	}
	return Solution(len(input))
}

package main

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	inputBytes := make([]byte, len(input))
	copy(inputBytes, input)
	for i := 0; i < 2; i++ {
		NextValid(&inputBytes)
	}
	return Solution(inputBytes)
}

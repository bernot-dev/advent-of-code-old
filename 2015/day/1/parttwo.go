package main

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	position := EntersBasement(string(input))
	return Solution(position)
}

// EntersBasement determines the position of the character that causes Santa to enter the basement
func EntersBasement(s string) int {
	currentFloor := 0
	for i, v := range s {
		switch v {
		case '(':
			currentFloor++
		case ')':
			currentFloor--
		}
		if currentFloor == -1 {
			return i + 1
		}
	}
	panic("Solution not found")
}

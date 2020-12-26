package main

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	floor := Floor(string(input))
	return Solution(floor)
}

// Floor determines the number of the floor, based on the input strings
func Floor(s string) (floor int) {
	for _, v := range s {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return
}

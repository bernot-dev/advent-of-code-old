package main

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	return Solution(FindHash(string(input), BeginsWithSixZeroes))
}

// BeginsWithSixZeroes .
func BeginsWithSixZeroes(hash *[16]byte) bool {
	if hash[0] != 0 || hash[1] != 0 || hash[2] != 0 {
		return false
	}
	return true
}

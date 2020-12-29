package main

import (
	"strconv"
	"strings"
)

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	for _, line := range input {
		solution += Solution(MemoryCountQuote(line) - CodeCount(line))
	}
	return
}

// MemoryCountQuote .
func MemoryCountQuote(s string) int {
	str := strconv.Quote(s)
	return strings.Count(str, "") - 1
}

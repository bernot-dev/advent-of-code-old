package main

import (
	"strconv"
	"strings"
)

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	for _, line := range input {
		solution += Solution(CodeCount(line) - MemoryCountUnquote(line))
	}
	return
}

// CodeCount .
func CodeCount(s string) int {
	return strings.Count(s, "") - 1
}

// MemoryCountUnquote .
func MemoryCountUnquote(s string) (count int) {
	str, err := strconv.Unquote(s)
	if err != nil {
		panic("Could not handle string: " + s)
	}
	return strings.Count(str, "") - 1
}

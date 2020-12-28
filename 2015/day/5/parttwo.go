package main

import (
	"fmt"
	"regexp"
)

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	rules := []Rule{
		RepeatedPair,
		ThirdWheel,
	}
	for _, line := range input {
		if IsNice(line, rules) {
			solution++
		}
	}
	return
}

// RepeatedPair .
func RepeatedPair(s string) bool {
	// pairs := make(map[string]bool)
	for i := 1; i < len(s); i++ {
		pair := s[i-1 : i+1]
		pairRegexp := regexp.MustCompile(fmt.Sprintf("(%s)", pair))
		matches := pairRegexp.FindAllStringIndex(s, -1)
		if len(matches) > 1 && matches[0][1] <= matches[len(matches)-1][0] {
			return true
		}
	}
	return false
}

// ThirdWheel .
func ThirdWheel(s string) (t bool) {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			return true
		}
	}
	return
}

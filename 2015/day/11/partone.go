package main

import "strings"

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	inputBytes := make([]byte, len(input))
	copy(inputBytes, input)
	return Solution(NextValid(&inputBytes))
}

// NextValid increments string until it finds the next string that matches all the rules
func NextValid(b *[]byte) string {
	lastIndex := len(*b) - 1
	for Increment(b, lastIndex); !matchesRules(b); Increment(b, lastIndex) {
	}
	return string(*b)
}

// Increment string according to rules
func Increment(b *[]byte, index int) error {
	if (*b)[index] < 'z' {
		(*b)[index]++
	} else {
		(*b)[index] = 'a'
		Increment(b, index-1)
	}
	return nil
}

func matchesRules(b *[]byte) bool {
	return hasStraight(b) && !containsBannedLetters(b) && containsTwoPair(b)
}

func hasStraight(b *[]byte) bool {
	for i := 3; i <= len(*b); i++ {
		if (*b)[i-3]+2 == (*b)[i-2]+1 && (*b)[i-2]+1 == (*b)[i-1] {
			return true
		}
	}
	return false
}

func containsBannedLetters(b *[]byte) bool {
	return strings.ContainsAny(string(*b), "iol")
}

func containsTwoPair(b *[]byte) bool {
	pairCount := 0
	firstMatch := byte(' ')
	for i := 1; i < len(*b); i++ {
		if (*b)[i-1] == (*b)[i] {
			if firstMatch != (*b)[i] {
				firstMatch = (*b)[i]
				pairCount++
			}
		}
	}
	return pairCount >= 2
}

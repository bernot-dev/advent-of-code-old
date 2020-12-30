package main

import "strconv"

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	for i := 0; i < 40; i++ {
		input = LookAndSay(input)
	}
	return Solution(len(input))
}

// LookAndSay transforms the input according to the rules of the game
func LookAndSay(s []byte) (result []byte) {
	result = make([]byte, 0)
	token := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			count++
			continue
		}
		result = append(result, []byte(strconv.Itoa(count))...)
		result = append(result, token)
		token = s[i]
		count = 1
	}
	result = append(result, []byte(strconv.Itoa(count))...)
	result = append(result, token)
	return
}

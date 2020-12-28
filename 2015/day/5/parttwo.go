package main

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
	for i := 1; i < len(s); i++ {
		for j := i + 2; j < len(s); j++ {
			if s[i-1] == s[j-1] && s[i] == s[j] {
				return true
			}
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

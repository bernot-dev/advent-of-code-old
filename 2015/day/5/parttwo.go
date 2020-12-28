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

// RepeatedPair determines if string contains a pair of any two letters that appears at least twice in
// the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
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

// ThirdWheel determines if string contains at least one letter which repeats with exactly one letter
//  between them, like xyx, abcdefeghi (efe), or even aaa.
func ThirdWheel(s string) (t bool) {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-2] {
			return true
		}
	}
	return
}

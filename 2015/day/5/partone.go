package main

import "regexp"

// Rule .
type Rule func(string) bool

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	rules := []Rule{
		ThreeVowels,
		DoubleLetter,
		NoBadPattern,
	}
	for _, line := range input {
		if IsNice(line, rules) {
			solution++
		}
	}
	return
}

// IsNice returns true if the string is nice
func IsNice(s string, rules []Rule) (nice bool) {
	nice = true
	for _, rule := range rules {
		nice = nice && rule(s)
	}
	return nice
}

// VowelCount .
func VowelCount(s string) (c int) {
	for _, ch := range s {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			c++
		}
	}
	return c
}

// ThreeVowels .
func ThreeVowels(s string) bool {
	return VowelCount(s) >= 3
}

// DoubleLetter .
func DoubleLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
	}
	return false
}

// NoBadPattern .
func NoBadPattern(s string) bool {
	badRegexp := regexp.MustCompile(`ab|cd|pq|xy`)
	return !badRegexp.MatchString(s)
}

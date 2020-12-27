package main

import (
	"crypto/md5"
	"strconv"
)

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	return Solution(FindHash(string(input), BeginsWithFiveZeroes))
}

// FindHash finds the lowest n-digit number that can be appended to the input string
// where the MD5 hash of the concatenated string begins with 5 zeroes
func FindHash(s string, f func(*[16]byte) bool) int {
	inputBytes := []byte(s)
	var numberBytes []byte

	var sum [16]byte
	for i := 0; true; i++ {
		numberBytes = []byte(strconv.Itoa(i))
		bytes := append(inputBytes, numberBytes...)
		sum = md5.Sum(bytes)
		if f(&sum) {
			return i
		}
	}
	panic("No hash found!")
}

// BeginsWithFiveZeroes .
func BeginsWithFiveZeroes(hash *[16]byte) bool {
	if hash[0] != 0 || hash[1] != 0 || hash[2]&0xF0 != 0 {
		return false
	}
	return true
}

package main

import "strconv"

func part1(in inputType) string {
	var count int
	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] {
			count++
		}
	}
	return strconv.Itoa(count)
}

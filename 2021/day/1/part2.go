package main

import "strconv"

func part2(in inputType) string {
	var count int

	var windowSums []int
	for i := 2; i < len(in); i++ {
		windowSums = append(windowSums, (in[i-2] + in[i-1] + in[i]))
	}

	for i := 1; i < len(windowSums); i++ {
		if windowSums[i] > windowSums[i-1] {
			count++
		}
	}
	return strconv.Itoa(count)
}

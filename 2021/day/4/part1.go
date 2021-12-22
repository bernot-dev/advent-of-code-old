package main

import "strconv"

func Part1(numbers []int, cards []*BingoCard) string {
	for _, number := range numbers {
		for _, card := range cards {
			bingo, score := card.Mark(number)
			if bingo {
				return strconv.Itoa(score)
			}
		}
	}
	return "fail"
}

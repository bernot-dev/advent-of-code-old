package main

import (
	"strconv"
)

func Part2(numbers []int, cards []*BingoCard) string {
	for _, number := range numbers {
		for i := 0; i < len(cards); i++ {
			bingo, score := cards[i].Mark(number)
			if bingo {
				cards[i] = cards[len(cards)-1]
				cards = cards[:len(cards)-1]
				if len(cards) == 0 {
					return strconv.Itoa(score)
				}
				i-- // iterate over newly swapped-in card
			}
		}
	}
	return "fail"
}

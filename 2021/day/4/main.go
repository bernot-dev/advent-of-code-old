package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2021, Day 4:")
	numbers, cards := readInput("input.txt")
	fmt.Printf("Part 1: %s\n", Part1(numbers, cards))
	fmt.Printf("Part 2: %s\n", Part2(numbers, cards))
}

type BingoCard struct {
	// 0-4 = rows, 5-9 = columns
	series []map[int]bool
	score  int
}

func NewBingoCard() *BingoCard {
	b := BingoCard{}
	b.series = make([]map[int]bool, 10, 10)
	for i := range b.series {
		b.series[i] = make(map[int]bool)
	}
	return &b
}

func (b *BingoCard) Mark(n int) (bingo bool, score int) {
	for i, line := range b.series {
		if line[n] {
			if i < 5 {
				b.score -= n
			}
			delete(b.series[i], n)
			if len(b.series[i]) == 0 {
				return true, n * b.score
			}
		}
	}
	return false, 0
}

type Bingo struct {
	remaining map[int]bool
	called    map[int]bool
}

func readInput(fileName string) (numbers []int, cards []*BingoCard) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	if scanner.Scan() {
		rawNumbers := scanner.Text()
		numbers = parseNumbers(rawNumbers)
	}

	for scanner.Scan() {
		card := NewBingoCard()
		card.parseBingoCard(scanner)
		cards = append(cards, card)
	}
	return numbers, cards
}

func parseNumbers(rawNumbers string) []int {
	numbers := make([]int, 0)
	strNumbers := strings.Split(rawNumbers, ",")

	for _, v := range strNumbers {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, value)
	}
	return numbers
}

func (b *BingoCard) parseBingoCard(scanner *bufio.Scanner) error {
	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := scanner.Text()
		numPattern := regexp.MustCompile(`\d+`)
		strNums := numPattern.FindAllString(line, -1)
		for j := 0; j < len(strNums); j++ {
			n, err := strconv.Atoi(strNums[j])
			if err != nil {
				return fmt.Errorf("converting %q to int: %w", strNums[j], err)
			}
			b.series[i][n] = true   // Add n to row
			b.series[5+j][n] = true // Add n to column
			b.score += n            // Sum numbers
		}
	}
	return nil
}

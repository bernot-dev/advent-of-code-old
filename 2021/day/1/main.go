package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	year = 2021
	day  = 1
)

type inputType []int

func main() {
	fmt.Printf("Solutions for Advent of Code %d, Day %d\n", year, day)

	in := readInput("input.txt")

	fmt.Printf("Part 1: %s\n", part1(in))
	fmt.Printf("Part 2: %s\n", part2(in))
}

func readInput(inputFile string) inputType {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var in inputType
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		in = append(in, n)
	}
	return in
}

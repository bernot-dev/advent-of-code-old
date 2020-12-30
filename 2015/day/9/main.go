package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/bernot-dev/advent-of-code/helpers"
)

// Input is the type of incoming data
type Input map[string]map[string]int

// Solution is the type of the eventual solution
type Solution int

func main() {
	input := ReadInput("input.txt")

	p1 := PartOne(input)
	fmt.Println("Part 1 Solution:", p1)

	p2 := PartTwo(input)
	fmt.Println("Part 2 Solution:", p2)
}

// ReadInput reads the input file and converts it to an appropriate format
func ReadInput(f string) Input {
	inputFile, err := os.Open(f)
	defer inputFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Process data
	m := make(map[string]map[string]int)

	lines := helpers.StringLines(inputFile)
	for _, line := range lines {
		distanceRegexp := regexp.MustCompile(`(?P<from>\w+) to (?P<to>\w+) = (?P<distance>\d+)`)
		const (
			full = iota
			from
			to
			distance
		)
		if !distanceRegexp.MatchString(line) {
			panic("Unhandled input: " + line)
		}
		matches := distanceRegexp.FindStringSubmatch(line)
		if _, ok := m[matches[from]]; !ok {
			m[matches[from]] = make(map[string]int)
		}
		if _, ok := m[matches[to]]; !ok {
			m[matches[to]] = make(map[string]int)
		}
		d, err := strconv.Atoi(matches[distance])
		if err != nil {
			panic("Unable to convert distance: " + matches[distance])
		}
		m[matches[from]][matches[to]] = d
		m[matches[to]][matches[from]] = d
	}
	return Input(m)
}

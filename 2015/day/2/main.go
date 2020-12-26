package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bernot-dev/advent-of-code/helpers"
)

// Present represents a box from the input
type Present struct {
	l, w, h int
}

// NewPresent creates a new Present
func NewPresent(s string) Present {
	dimensions := strings.Split(s, "x")
	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])
	return Present{l: l, w: w, h: h}
}

// Input is the type of incoming data
type Input []Present

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
	lines := helpers.StringLines(inputFile)
	p := make([]Present, len(lines))
	for i, line := range lines {
		p[i] = NewPresent(line)
	}
	return p
}

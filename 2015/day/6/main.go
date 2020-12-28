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
type Input []Instruction

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
	input := helpers.StringLines(inputFile)
	instructions := make([]Instruction, len(input))
	for i, v := range input {
		instructions[i] = New(v)
	}
	return instructions
}

// New turns a string into an Instruction
func New(s string) Instruction {
	re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	match := re.FindAllStringSubmatch(s, 1)[0]
	command := match[1]
	startRow, _ := strconv.Atoi(match[2])
	startCol, _ := strconv.Atoi(match[3])
	endRow, _ := strconv.Atoi(match[4])
	endCol, _ := strconv.Atoi(match[5])

	return Instruction{
		command: command,
		start:   Coordinate{startRow, startCol},
		end:     Coordinate{endRow, endCol},
	}
}

// Coordinate .
type Coordinate struct {
	x, y int
}

// Instruction .
type Instruction struct {
	command string
	start   Coordinate
	end     Coordinate
}

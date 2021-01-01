package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bernot-dev/advent-of-code/helpers"
)

// Input is the type of incoming data
type Input string

// Solution is the type of the eventual solution
type Solution string

func main() {
	p1 := PartOne(ReadInput("input.txt"))
	fmt.Println("Part 1 Solution:", p1)

	p2 := PartTwo(ReadInput("input.txt"))
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
	return Input(helpers.String(inputFile))
}

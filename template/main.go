package main

import (
	"fmt"
	"log"
	"os"
)

// Input is the type of incoming data
type Input []string

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
	return nil
}

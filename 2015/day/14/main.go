package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Reindeer .
type Reindeer struct {
	Name     string
	Speed    int
	Duration int
	Rest     int
}

// Input is the type of incoming data
type Input []Reindeer

// Solution is the type of the eventual solution
type Solution int

func main() {
	p1 := PartOne(ReadInput("input.txt"), 2503)
	fmt.Println("Part 1 Solution:", p1)

	p2 := PartTwo(ReadInput("input.txt"), 2503)
	fmt.Println("Part 2 Solution:", p2)
}

// ReadInput reads the input file and converts it to an appropriate format
func ReadInput(f string) Input {
	inputFile, err := os.Open(f)
	defer inputFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

	input := make([]Reindeer, 0)

	// Process data
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		reindeer := Reindeer{
			Name:     matches[1],
			Speed:    mustInt(matches[2]),
			Duration: mustInt(matches[3]),
			Rest:     mustInt(matches[4]),
		}
		input = append(input, reindeer)
	}
	return Input(input)
}

func mustInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

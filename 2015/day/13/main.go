package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Edge .
type Edge struct {
	A string
	B string
}

// Graph .
type Graph map[Edge]int

// NewGraph .
func NewGraph() Graph {
	return make(Graph)
}

// Input is the type of incoming data
type Input struct {
	Graph Graph
	Names []string
	First string
}

// Solution is the type of the eventual solution
type Solution int

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

	re := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)`)

	m := make(map[Edge]int)
	names := make(map[string]bool)

	// Process data
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		name, sign, pointsStr, neighbor := matches[1], matches[2], matches[3], matches[4]
		if name == "" || neighbor == "" {
			log.Fatal("empty name")
		}
		e := Edge{name, neighbor}
		weight := pointsToInt(pointsStr, sign)

		names[name] = true
		names[neighbor] = true
		m[e] = weight
	}

	nameList := make([]string, 0, len(names))
	for name := range names {
		nameList = append(nameList, name)
	}

	return Input{
		First: nameList[0],
		Graph: m,
		Names: nameList,
	}
}

func pointsToInt(pointsStr, sign string) int {
	points, err := strconv.ParseInt(pointsStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if sign == "lose" {
		return int(points) * -1
	}
	return int(points)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Sue .
type Sue struct {
	N           int
	Children    int
	Cats        int
	Samoyeds    int
	Pomeranians int
	Akitas      int
	Vizslas     int
	Goldfish    int
	Trees       int
	Cars        int
	Perfumes    int
}

// Input is the type of incoming data
type Input []Sue

// Solution is the type of the eventual solution
type Solution int

func main() {
	p1 := PartOne(ReadInput("input.txt"))
	fmt.Println("Part 1 Solution:", p1)

	p2 := PartTwo(ReadInput("input.txt"))
	fmt.Println("Part 2 Solution:", p2)
}

// Sue 1: goldfish: 9, cars: 0, samoyeds: 9

// ReadInput reads the input file and converts it to an appropriate format
func ReadInput(f string) Input {
	inputFile, err := os.Open(f)
	defer inputFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	suere := regexp.MustCompile(`Sue (\d+):`)
	re := regexp.MustCompile(`(\w+): (\d+)`)

	sues := make([]Sue, 0)

	// Process data
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		sue := suere.FindStringSubmatch(line)[1]
		parts := re.FindAllStringSubmatch(line, -1)
		// fmt.Printf("Sue: %s, %#v\n", sue, parts)

		s := Sue{
			N:           mustInt(sue),
			Children:    -1,
			Cats:        -1,
			Samoyeds:    -1,
			Pomeranians: -1,
			Akitas:      -1,
			Vizslas:     -1,
			Goldfish:    -1,
			Trees:       -1,
			Cars:        -1,
			Perfumes:    -1,
		}
		for i, part := range parts {
			if i == -1 {
				continue
			}
			switch part[1] {
			case "children":
				s.Children = mustInt(part[2])
			case "cats":
				s.Cats = mustInt(part[2])
			case "samoyeds":
				s.Samoyeds = mustInt(part[2])
			case "pomeranians":
				s.Pomeranians = mustInt(part[2])
			case "akitas":
				s.Akitas = mustInt(part[2])
			case "vizslas":
				s.Vizslas = mustInt(part[2])
			case "goldfish":
				s.Goldfish = mustInt(part[2])
			case "trees":
				s.Trees = mustInt(part[2])
			case "cars":
				s.Cars = mustInt(part[2])
			case "perfumes":
				s.Perfumes = mustInt(part[2])
			default:
				log.Fatalf("unhandled thing: %s", part[1])

			}
		}
		sues = append(sues, s)
	}
	return sues
}

func mustInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

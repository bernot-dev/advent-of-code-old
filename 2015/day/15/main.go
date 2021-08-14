package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Ingredient is the type of incoming data
type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
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
func ReadInput(f string) []Ingredient {
	inputFile, err := os.Open(f)
	defer inputFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (\d+)`)
	ingredients := make([]Ingredient, 0)

	scanner := bufio.NewScanner(inputFile)
	// Process data
	for scanner.Scan() {
		line := scanner.Text()
		parts := re.FindStringSubmatch(line)
		ingredients = append(ingredients, Ingredient{
			Name:       parts[1],
			Capacity:   mustInt(parts[2]),
			Durability: mustInt(parts[3]),
			Flavor:     mustInt(parts[4]),
			Texture:    mustInt(parts[5]),
			Calories:   mustInt(parts[6]),
		})
	}
	return ingredients
}

func mustInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

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
type Input []Gate

// Solution is the type of the eventual solution
type Solution uint16

func main() {
	input := ReadInput("input.txt")

	p1 := PartOne(input, "a")
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
	gates := make([]Gate, len(lines))
	for i, line := range lines {
		gates[i] = NewGate(line)
	}
	return gates
}

// Wire .
type Wire struct {
	id string
	c  chan uint16
}

// Gate .
type Gate struct {
	raw          string
	operator     func(...uint16) uint16
	inputWireIDs []interface{}
	outputWireID string

	inputWires []Wire
	outputWire Wire
}

// NewGate decodes a string and turns it into a gate
func NewGate(s string) (g Gate) {
	gateRegexp := regexp.MustCompile(`^(?P<left>[a-z0-9]+)?? ?(?P<operator>[A-Z]+)? ?(?P<right>[a-z0-9]+) -> (?P<out>[a-z]+)$`)
	const ( // Named capture fields from regex
		raw = iota
		left
		operator
		right
		out
	)
	if !gateRegexp.MatchString(s) {
		panic("Unhandled instruction: " + s)
	}
	match := gateRegexp.FindStringSubmatch(s)
	g.raw = match[raw]
	if match[operator] == "" {
		match[operator] = "SIGNAL"
	}
	g.operator = ops[match[operator]]
	g.inputWireIDs = make([]interface{}, 0)
	g.outputWireID = match[out]
	for _, operand := range []string{match[left], match[right]} {
		if operand != "" {
			if n, err := strconv.Atoi(operand); err == nil {
				g.inputWireIDs = append(g.inputWireIDs, uint16(n))
			} else {
				g.inputWireIDs = append(g.inputWireIDs, operand)
			}
		}
	}
	if g.outputWireID == "" {
		panic(g.raw)
	}
	return
}

// Operator .
type Operator func(...uint16) uint16

var ops = map[string]Operator{
	"SIGNAL": func(n ...uint16) uint16 { return n[0] },
	"AND":    func(n ...uint16) uint16 { return n[0] & n[1] },
	"OR":     func(n ...uint16) uint16 { return n[0] | n[1] },
	"NOT":    func(n ...uint16) uint16 { return ^n[0] },
	"LSHIFT": func(n ...uint16) uint16 { return n[0] << n[1] },
	"RSHIFT": func(n ...uint16) uint16 { return n[0] >> n[1] },
}

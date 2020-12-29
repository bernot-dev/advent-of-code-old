package main

import (
	"fmt"
	"strconv"
)

// Circuit .
type Circuit struct {
	wires map[string]Wire
	gates []Gate
}

// PartOne solves the first part of the problem.
func PartOne(input Input, solutionWire string) (solution Solution) {
	c := Circuit{
		gates: []Gate(input),
		wires: nil,
	}
	c.CreateWires()
	c.ConnectGates()
	return Solution(<-c.wires[solutionWire].c)
}

// CreateWires sets up a map of all wires
func (c *Circuit) CreateWires() {
	c.wires = make(map[string]Wire)
	for _, gate := range c.gates {
		// Test every wire connected to the gate, input and output
		for _, wireDescriptor := range []interface{}{gate.inputWireIDs, gate.outputWireID} {
			// Skip numeric wires - they will be created ad-hoc
			if wireID, ok := wireDescriptor.(string); ok {
				if _, ok := c.wires[wireID]; !ok {
					c.wires[wireID] = Wire{
						id: wireID,
						c:  make(chan uint16),
					}
				}
			}
		}
	}
}

// ConnectGates connects the gates using the wires
func (c *Circuit) ConnectGates() {
	for i := range c.gates {
		for _, inputWire := range c.gates[i].inputWireIDs {
			c.gates[i].inputWires = append(c.gates[i].inputWires, c.Resolve(inputWire))
		}
		c.gates[i].outputWire = c.wires[c.gates[i].outputWireID]

		go activate(c.gates[i].inputWires, c.gates[i].outputWire, c.gates[i].operator)

	}
}

// Resolve returns the appropriate wire, based on the descriptor, generating a literal signal if necessary
func (c *Circuit) Resolve(i interface{}) Wire {
	// Connect an existing wire
	if s, ok := i.(string); ok {
		return c.wires[s]
	}

	// Connect a new wire that provides the specified signal value
	if n, ok := i.(uint16); ok {
		return Wire{
			id: strconv.Itoa(int(n)),
			c:  signal(n),
		}
	}
	panic(fmt.Sprintf("Unable to decode value: %#v", i))
}

// signal creates a channel that repeats a given value
func signal(n uint16) chan uint16 {
	ch := make(chan uint16, 10)
	go func(n uint16, out chan uint16) {
		for {
			out <- n
		}
	}(n, ch)
	return ch
}

// activate makes signal start flowing through the switch
func activate(in []Wire, out Wire, operator Operator) {
	input := make([]uint16, len(in))
	// Receive input signals
	for i, v := range in {
		input[i] = <-v.c
	}
	// Transform input to result
	result := operator(input...)

	// Broadcast result
	for {
		out.c <- result
	}
}

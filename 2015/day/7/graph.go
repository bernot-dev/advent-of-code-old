package main

import (
	"fmt"
	"io"
	"regexp"
)

var (
	numeric = regexp.MustCompile(`^\d+$`)
)

func graph(in Input, output io.WriteCloser) {
	c := Circuit{
		gates: []Gate(in),
		wires: nil,
	}
	c.CreateWires()
	c.ConnectGates()
	fmt.Fprintln(output, "digraph G {")
	for _, gate := range c.gates {
		for _, wire := range gate.inputWires {
			if !numeric.MatchString(wire.id) {
				label := fmt.Sprintf(" %s:%d", wire.id, <-wire.c)
				fmt.Fprintf(output, "%s -> %s [label=%q]\n", wire.id, gate.outputWire.id, label)
			}
		}
	}
	for _, gate := range c.gates {
		label := fmt.Sprintf("%s [%d]", gate.raw, <-gate.outputWire.c)
		fmt.Fprintf(output, "%s [label=%q]\n", gate.outputWire.id, label)
	}
	fmt.Fprintf(output, "a -> solution [label=%q]\n", "a: 16076")
	fmt.Fprintln(output, "}")
}

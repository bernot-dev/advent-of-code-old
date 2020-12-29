package main

// PartTwo sovles the second part of the problem.
func PartTwo(input Input, solutionWire string) (solution Solution) {
	revisedInput := make(Input, 0)
	for _, ins := range input {
		if ins.outputWireID == "b" {
			revisedInput = append(revisedInput, NewGate("16076 -> b"))
		} else {
			revisedInput = append(revisedInput, ins)
		}
	}
	c := Circuit{
		gates: []Gate(revisedInput),
		wires: nil,
	}
	c.CreateWires()
	c.ConnectGates()
	return Solution(<-c.wires[solutionWire].c)
}

package main

// PartTwo sovles the second part of the problem.
func PartTwo(input Input) (solution Solution) {
	m := make(map[Vector]int)
	santa := Vector{}
	robo := Vector{}
	m[santa]++
	m[robo]++
	for i, v := range input {
		switch i % 2 {
		case 0:
			switch v {
			case '^':
				santa.Add(N)
			case '>':
				santa.Add(E)
			case 'v':
				santa.Add(S)
			case '<':
				santa.Add(W)
			}
			m[santa]++
		case 1:
			switch v {
			case '^':
				robo.Add(N)
			case '>':
				robo.Add(E)
			case 'v':
				robo.Add(S)
			case '<':
				robo.Add(W)
			}
			m[robo]++
		}
	}
	return Solution(len(m))
}

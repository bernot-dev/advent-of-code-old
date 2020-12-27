package main

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	m := make(map[Vector]int)
	cur := Vector{}
	m[cur]++
	for _, v := range input {
		switch v {
		case '^':
			cur.Add(N)
		case '>':
			cur.Add(E)
		case 'v':
			cur.Add(S)
		case '<':
			cur.Add(W)
		}
		m[cur]++
	}
	return Solution(len(m))
}

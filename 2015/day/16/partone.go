package main

// Equals .
func (s Sue) Equals(a Sue) bool {
	if a.Children >= 0 && s.Children >= 0 && a.Children != s.Children {
		return false
	}
	if a.Cats >= 0 && s.Cats >= 0 && a.Cats != s.Cats {
		return false
	}
	if a.Samoyeds >= 0 && s.Samoyeds >= 0 && a.Samoyeds != s.Samoyeds {
		return false
	}
	if a.Pomeranians >= 0 && s.Pomeranians >= 0 && a.Pomeranians != s.Pomeranians {
		return false
	}
	if a.Akitas >= 0 && s.Akitas >= 0 && a.Akitas != s.Akitas {
		return false
	}
	if a.Vizslas >= 0 && s.Vizslas >= 0 && a.Vizslas != s.Vizslas {
		return false
	}
	if a.Goldfish >= 0 && s.Goldfish >= 0 && a.Goldfish != s.Goldfish {
		return false
	}
	if a.Trees >= 0 && s.Trees >= 0 && a.Trees != s.Trees {
		return false
	}
	if a.Cars >= 0 && s.Cars >= 0 && a.Cars != s.Cars {
		return false
	}
	if a.Perfumes >= 0 && s.Perfumes >= 0 && a.Perfumes != s.Perfumes {
		return false
	}
	return true
}

// PartOne solves the first part of the problem.
func PartOne(input Input) (solution Solution) {
	gifter := Sue{
		Children:    3,
		Cats:        7,
		Samoyeds:    2,
		Pomeranians: 3,
		Akitas:      0,
		Vizslas:     0,
		Goldfish:    5,
		Trees:       3,
		Cars:        2,
		Perfumes:    1,
	}
	for _, v := range input {
		if gifter.Equals(v) {
			return Solution(v.N)
		}
	}
	return
}

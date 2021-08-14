package main

import (
	"math"
)

// Calories .
func (r Recipe) Calories() int {
	c := 0
	for i, v := range r.Ingredients {
		c += v.Calories * r.Quantities[i]
	}
	return c
}

// PartTwo sovles the second part of the problem.
func PartTwo(input []Ingredient) (solution Solution) {
	max := math.MinInt64
	for i := 0; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			for k := 0; k < 100-j-i; k++ {

				last := 100 - i - j - k
				r := NewRecipe(input, []int{i, j, k, last})
				if r.Calories() != 500 {
					continue
				}
				score := r.Score()
				if score > max {
					max = score
				}
			}
		}
	}
	return Solution(max)

}

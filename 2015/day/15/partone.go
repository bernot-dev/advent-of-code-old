package main

import (
	"log"
	"math"
)

// Recipe .
type Recipe struct {
	Ingredients []Ingredient
	Quantities  []int
}

// NewRecipe .
func NewRecipe(ingredients []Ingredient, quantities []int) Recipe {
	if len(ingredients) != len(quantities) {
		log.Fatal("ingredients do not match quantities")
	}

	return Recipe{
		Ingredients: ingredients,
		Quantities:  quantities,
	}
}

// Score .
func (r Recipe) Score() int {
	final := Ingredient{}
	for i, v := range r.Ingredients {
		final.Capacity += v.Capacity * r.Quantities[i]
		final.Durability += v.Durability * r.Quantities[i]
		final.Flavor += v.Flavor * r.Quantities[i]
		final.Texture += v.Texture * r.Quantities[i]
	}

	if final.Capacity < 0 {
		final.Capacity = 0
	}

	if final.Durability < 0 {
		final.Durability = 0
	}

	if final.Flavor < 0 {
		final.Flavor = 0
	}

	if final.Texture < 0 {
		final.Texture = 0
	}

	return final.Capacity * final.Durability * final.Flavor * final.Texture
}

// Add .
func (i Ingredient) Add(a Ingredient) Ingredient {
	return Ingredient{
		Capacity:   i.Capacity + a.Capacity,
		Durability: i.Durability + a.Durability,
		Flavor:     i.Flavor + a.Flavor,
		Texture:    i.Texture + a.Texture,
		Calories:   i.Calories + a.Calories,
	}
}

// Multiply .
func (i Ingredient) Multiply(n int) Ingredient {
	return Ingredient{
		Capacity:   i.Capacity * n,
		Durability: i.Durability * n,
		Flavor:     i.Flavor * n,
		Texture:    i.Texture * n,
		Calories:   i.Calories * n,
	}
}

// PartOne solves the first part of the problem.
func PartOne(input []Ingredient) (solution Solution) {
	max := math.MinInt64
	for i := 0; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			for k := 0; k < 100-j-i; k++ {

				last := 100 - i - j - k
				{
					r := NewRecipe(input, []int{i, j, k, last})
					score := r.Score()
					if score > max {
						max = score
					}
				}
			}
		}
	}
	return Solution(max)
}

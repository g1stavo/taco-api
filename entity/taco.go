package entity

import "strings"

// Ingredient is a taco ingredient
type Ingredient struct {
	URL        string `json:"url"`
	Name       string `json:"name"`
	Vegan      bool   `json:"vegan"`
	Vegetarian bool   `json:"vegetarian"`
}

// SetVegFields set ingredient veg flags
func (i *Ingredient) SetVegFields(description string) {
	i.Vegan = strings.Contains(description, "vegan")
	i.Vegetarian = strings.Contains(description, "vegetarian")
}

// Taco is a taco
type Taco struct {
	Seasoning Ingredient `json:"seasoning"`
	Condiment Ingredient `json:"condiment"`
	Mixin     Ingredient `json:"mixin"`
	BaseLayer Ingredient `json:"base_layer"`
	Shell     Ingredient `json:"shell"`
}

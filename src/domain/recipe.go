package domain

type RecipeName string

type Recipe struct {
	Name        RecipeName
	Ingredients []IngredientName
}

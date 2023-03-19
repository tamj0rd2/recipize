package random

import (
	"fmt"
	"github.com/tamj0rd2/recipize/src/domain"
	"time"
)

func Recipe() domain.Recipe {
	return domain.Recipe{
		Name:        RecipeName(),
		Ingredients: Ingredients(),
	}
}

func RecipeName() domain.RecipeName {
	return domain.RecipeName(fmt.Sprintf("My Recipe %s", time.Now().UTC().String()))
}

func Ingredients() []domain.IngredientName {
	options := []domain.IngredientName{"Potato", "Tomato", "Carrot", "Pepper", "Chicken breast", "Macaroni"}
	return options[:Int(1, len(options))]
}

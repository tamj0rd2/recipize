package domainspecs

import (
	"context"
	"github.com/tamj0rd2/recipize/src/domain"
	"testing"
)

type RecipizeDSL interface {
	GivenTheRecipeDoesNotExist(name domain.RecipeName)
	WhenIAttemptToCreateTheRecipe(recipe domain.Recipe)
	ThenICanSeeTheRecipeInMyRecipesList(name domain.RecipeName)
	ThenICanSeeTheIngredientsOfTheRecipe(name domain.RecipeName, ingredients []domain.IngredientName)
}

func TestRecipize(t *testing.T, makeDriver func(ctx context.Context, t testing.TB) RecipizeDSL) {
	t.Run("Creating a recipe", func(t *testing.T) {
		driver := makeDriver(context.Background(), t)

		recipe := domain.Recipe{
			Name:        "Scrambled eggs",
			Ingredients: []domain.IngredientName{"Egg", "Butter", "Salt", "Pepper", "Onion"},
		}

		driver.GivenTheRecipeDoesNotExist(recipe.Name)
		driver.WhenIAttemptToCreateTheRecipe(recipe)
		driver.ThenICanSeeTheRecipeInMyRecipesList(recipe.Name)
		driver.ThenICanSeeTheIngredientsOfTheRecipe(recipe.Name, recipe.Ingredients)
	})
}

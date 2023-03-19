package domainspecs

import (
	"context"
	"github.com/tamj0rd2/recipize/src/domain"
	"testing"
)

type RecipizeDSL interface {
	GivenTheRecipeDoesNotExist(name domain.RecipeName)
	WhenIAttemptToCreateTheRecipe(name domain.RecipeName, ingredients []domain.IngredientName)
	ThenICanSeeTheRecipeInMyRecipesList(name domain.RecipeName)
	ThenICanSeeTheIngredientsOfTheRecipe(name domain.RecipeName, ingredients []domain.IngredientName)
}

func TestRecipize(t *testing.T, makeDriver func(ctx context.Context, t testing.TB) RecipizeDSL) {
	t.Run("Creating a recipe", func(t *testing.T) {
		driver := makeDriver(context.Background(), t)

		recipeName := domain.RecipeName("Scrambled eggs")
		ingredients := []domain.IngredientName{"Egg", "Butter", "Salt", "Pepper", "Onion"}

		driver.GivenTheRecipeDoesNotExist(recipeName)
		driver.WhenIAttemptToCreateTheRecipe(recipeName, ingredients)
		driver.ThenICanSeeTheRecipeInMyRecipesList(recipeName)
		driver.ThenICanSeeTheIngredientsOfTheRecipe(recipeName, ingredients)
	})
}

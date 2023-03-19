package domainspecs

import (
	"context"
	"github.com/tamj0rd2/recipize/src/domain"
	"testing"
)

type RecipizeDSL interface {
	GivenTheRecipeDoesNotExist(name domain.RecipeName)
	GivenIHaveCreatedARecipe() domain.Recipe
	WhenITryToCreateTheRecipe(recipe domain.Recipe)
	WhenIAddTheRecipeToMyMealPlan(name domain.RecipeName)
	ThenICanSeeTheRecipeInMyList(name domain.RecipeName)
	ThenICanSeeTheIngredientsOfTheRecipe(name domain.RecipeName, ingredients []domain.IngredientName)
	ThenICanSeeTheRecipeInMyMealPlan(name domain.RecipeName)
	GivenTheRecipeIsNotInMyMealPlan(name domain.RecipeName)
}

func TestRecipize(t *testing.T, makeDriver func(ctx context.Context, t testing.TB) RecipizeDSL) {
	t.Run("Creating a recipe", func(t *testing.T) {
		driver := makeDriver(context.Background(), t)

		recipe := domain.Recipe{
			Name:        "Scrambled eggs",
			Ingredients: []domain.IngredientName{"Egg", "Butter", "Salt", "Pepper", "Onion"},
		}

		driver.GivenTheRecipeDoesNotExist(recipe.Name)
		driver.WhenITryToCreateTheRecipe(recipe)
		driver.ThenICanSeeTheRecipeInMyList(recipe.Name)
		driver.ThenICanSeeTheIngredientsOfTheRecipe(recipe.Name, recipe.Ingredients)
	})

	t.Run("Adding a recipe to my meal plan", func(t *testing.T) {
		driver := makeDriver(context.Background(), t)

		recipe := driver.GivenIHaveCreatedARecipe()
		driver.GivenTheRecipeIsNotInMyMealPlan(recipe.Name)
		driver.WhenIAddTheRecipeToMyMealPlan(recipe.Name)
		driver.ThenICanSeeTheRecipeInMyMealPlan(recipe.Name)
	})
}

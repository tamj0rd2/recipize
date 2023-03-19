package domainspecs

import (
	"context"
	"github.com/tamj0rd2/recipize/src/domain"
	"testing"
)

type RecipizeDSL interface {
	GivenTheRecipeDoesNotExist(name domain.RecipeName)
	WhenIAttemptToCreateTheRecipe(name domain.RecipeName)
	ThenICanSeeTheRecipeInMyRecipesList(name domain.RecipeName)
}

func TestRecipize(t *testing.T, makeDriver func(ctx context.Context, t testing.TB) RecipizeDSL) {
	t.Run("Creating a recipe", func(t *testing.T) {
		driver := makeDriver(context.Background(), t)

		recipeName := domain.RecipeName("Scrambled eggs")
		driver.GivenTheRecipeDoesNotExist(recipeName)
		driver.WhenIAttemptToCreateTheRecipe(recipeName)
		driver.ThenICanSeeTheRecipeInMyRecipesList(recipeName)
	})
}

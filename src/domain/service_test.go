package domain_test

import (
	"context"
	"github.com/tamj0rd2/recipize/src/adapters/memory"
	"github.com/tamj0rd2/recipize/src/domain"
	"github.com/tamj0rd2/recipize/src/domain/domainspecs"
	"testing"
)

func TestService(t *testing.T) {
	recipeStorage := memory.NewRecipeStorage()

	domainspecs.TestRecipize(t, func(ctx context.Context, t testing.TB) domainspecs.RecipizeDSL {
		return domainspecs.NewDriver(ctx, t, domain.NewService(recipeStorage))
	})
}

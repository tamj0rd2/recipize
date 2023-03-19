package ports

import (
	"context"
	"github.com/tamj0rd2/recipize/src/domain"
)

type Recipize interface {
	DoesRecipeExist(ctx context.Context, name domain.RecipeName) (bool, error)
	CreateRecipe(ctx context.Context, name domain.RecipeName) error
	GetRecipes(ctx context.Context) ([]domain.RecipeName, error)
}

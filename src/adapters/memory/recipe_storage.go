package memory

import (
	"context"
	"github.com/tamj0rd2/recipize/src/domain"
)

func NewRecipeStorage() *RecipeStorage {
	return &RecipeStorage{
		recipes: make(map[domain.RecipeName]bool),
	}
}

type RecipeStorage struct {
	recipes map[domain.RecipeName]bool
}

func (r *RecipeStorage) GetRecipes(ctx context.Context) ([]domain.RecipeName, error) {
	recipes := make([]domain.RecipeName, 0, len(r.recipes))
	for recipe := range r.recipes {
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (r *RecipeStorage) DoesRecipeExist(ctx context.Context, name domain.RecipeName) (bool, error) {
	_, ok := r.recipes[name]
	return ok, nil
}

func (r *RecipeStorage) CreateRecipe(ctx context.Context, name domain.RecipeName) error {
	r.recipes[name] = true
	return nil
}

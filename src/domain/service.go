package domain

import (
	"context"
	"fmt"
)

type RecipeStorage interface {
	CreateRecipe(ctx context.Context, name RecipeName) error
	DoesRecipeExist(ctx context.Context, name RecipeName) (bool, error)
	GetRecipes(ctx context.Context) ([]RecipeName, error)
}

func NewService(recipeStorage RecipeStorage) *Service {
	return &Service{recipeStorage: recipeStorage}
}

type Service struct {
	recipeStorage RecipeStorage
}

func (s Service) DoesRecipeExist(ctx context.Context, name RecipeName) (bool, error) {
	doesExist, err := s.recipeStorage.DoesRecipeExist(ctx, name)
	if err != nil {
		return false, fmt.Errorf("failed to check if recipe exists: %w", err)
	}

	return doesExist, nil
}

func (s Service) CreateRecipe(ctx context.Context, name RecipeName) error {
	if err := s.recipeStorage.CreateRecipe(ctx, name); err != nil {
		return fmt.Errorf("failed to create recipe: %w", err)
	}

	return nil
}

func (s Service) GetRecipes(ctx context.Context) ([]RecipeName, error) {
	recipes, err := s.recipeStorage.GetRecipes(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get recipes: %w", err)
	}

	return recipes, nil
}

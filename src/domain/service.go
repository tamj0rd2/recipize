package domain

import (
	"context"
	"fmt"
)

type RecipeStorage interface {
	CreateRecipe(ctx context.Context, recipe Recipe) error
	GetRecipes(ctx context.Context) ([]RecipeName, error)
	GetRecipe(ctx context.Context, name RecipeName) (Recipe, bool, error)
	DeleteRecipe(ctx context.Context, name RecipeName) error
	AddRecipeToMealPlan(ctx context.Context, recipe Recipe) error
	GetMealPlan(ctx context.Context) (MealPlan, error)
}

func NewService(recipeStorage RecipeStorage) *Service {
	return &Service{recipeStorage: recipeStorage}
}

type Service struct {
	recipeStorage RecipeStorage
}

func (s Service) GetRecipe(ctx context.Context, name RecipeName) (Recipe, bool, error) {
	recipe, found, err := s.recipeStorage.GetRecipe(ctx, name)
	if err != nil {
		return Recipe{}, false, fmt.Errorf("failed to get recipe: %w", err)
	}

	return recipe, found, nil
}

func (s Service) DoesRecipeExist(ctx context.Context, name RecipeName) (bool, error) {
	_, doesExist, err := s.recipeStorage.GetRecipe(ctx, name)
	if err != nil {
		return false, fmt.Errorf("failed to check if recipe exists: %w", err)
	}

	return doesExist, nil
}

func (s Service) CreateRecipe(ctx context.Context, recipe Recipe) error {
	if err := s.recipeStorage.CreateRecipe(ctx, recipe); err != nil {
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

func (s Service) AddRecipeToMealPlan(ctx context.Context, name RecipeName) error {
	if err := s.recipeStorage.AddRecipeToMealPlan(ctx, Recipe{Name: name}); err != nil {
		return fmt.Errorf("failed to add recipe to meal plan: %w", err)
	}

	return nil
}

func (s Service) GetMealPlan(ctx context.Context) (MealPlan, error) {
	mealPlan, err := s.recipeStorage.GetMealPlan(ctx)
	if err != nil {
		return MealPlan{}, fmt.Errorf("failed to get meal plan: %w", err)
	}

	return mealPlan, nil
}

func (s Service) DeleteRecipe(ctx context.Context, name RecipeName) error {
	if err := s.recipeStorage.DeleteRecipe(ctx, name); err != nil {
		return fmt.Errorf("failed to delete recipe: %w", err)
	}

	return nil
}

package memory

import (
	"context"
	"fmt"
	"github.com/tamj0rd2/recipize/lib/slices"
	"github.com/tamj0rd2/recipize/src/domain"
)

func NewRecipeStorage() *RecipeStorage {
	return &RecipeStorage{
		recipes: make(map[domain.RecipeName]domain.Recipe),
	}
}

type RecipeStorage struct {
	recipes  map[domain.RecipeName]domain.Recipe
	mealPlan []domain.RecipeName
}

func (r *RecipeStorage) DeleteRecipe(ctx context.Context, name domain.RecipeName) error {
	delete(r.recipes, name)
	return nil
}

func (r *RecipeStorage) GetRecipes(ctx context.Context) ([]domain.RecipeName, error) {
	recipes := make([]domain.RecipeName, 0, len(r.recipes))
	for recipe := range r.recipes {
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (r *RecipeStorage) GetRecipe(ctx context.Context, name domain.RecipeName) (domain.Recipe, bool, error) {
	recipe, ok := r.recipes[name]
	return recipe, ok, nil
}

func (r *RecipeStorage) CreateRecipe(ctx context.Context, recipe domain.Recipe) error {
	r.recipes[recipe.Name] = recipe
	return nil
}

func (r *RecipeStorage) GetMealPlan(ctx context.Context) (domain.MealPlan, error) {
	return domain.MealPlan{Recipes: r.mealPlan}, nil
}

func (r *RecipeStorage) AddRecipeToMealPlan(ctx context.Context, recipe domain.Recipe) error {
	_, ok := r.recipes[recipe.Name]
	if !ok {
		return fmt.Errorf("recipe %q does not exist", recipe.Name)
	}

	if slices.Contains(r.mealPlan, recipe.Name) {
		return fmt.Errorf("recipe %q is already in the meal plan", recipe.Name)
	}

	r.mealPlan = append(r.mealPlan, recipe.Name)
	return nil
}

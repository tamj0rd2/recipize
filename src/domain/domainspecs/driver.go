package domainspecs

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/tamj0rd2/recipize/lib/slices"
	"github.com/tamj0rd2/recipize/src/domain"
	"github.com/tamj0rd2/recipize/testhelpers/random"
	"testing"
)

type Recipize interface {
	CreateRecipe(ctx context.Context, recipe domain.Recipe) error
	GetRecipes(ctx context.Context) ([]domain.RecipeName, error)
	GetRecipe(ctx context.Context, name domain.RecipeName) (domain.Recipe, bool, error)
	DeleteRecipe(ctx context.Context, name domain.RecipeName) error
	AddRecipeToMealPlan(ctx context.Context, name domain.RecipeName) error
	GetMealPlan(ctx context.Context) (domain.MealPlan, error)
}

func NewDriver(ctx context.Context, t testing.TB, recipize Recipize) *Driver {
	return &Driver{
		ctx:      ctx,
		t:        t,
		Recipize: recipize,
	}
}

type Driver struct {
	t testing.TB
	Recipize
	ctx context.Context
}

func (d *Driver) GivenTheRecipeDoesNotExist(name domain.RecipeName) {
	t := d.t
	t.Helper()

	_, exists, err := d.GetRecipe(d.ctx, name)
	assert.NoError(t, err, "failed to check if recipe exists")
	assert.False(t, exists, "the recipe already exists")
}

func (d *Driver) GivenIHaveCreatedARecipe() domain.Recipe {
	recipe := random.Recipe()
	d.GivenTheRecipeDoesNotExist(recipe.Name)
	d.WhenITryToCreateTheRecipe(recipe)
	return recipe
}

func (d *Driver) GivenTheRecipeIsNotInMyMealPlan(name domain.RecipeName) {
	t := d.t
	t.Helper()

	mealPlan, err := d.Recipize.GetMealPlan(d.ctx)
	assert.NoError(t, err, "failed to get recipes")

	if slices.Contains(mealPlan.Recipes, name) {
		t.Fatalf("%q is already in the meal plan: %v", name, mealPlan)
	}
}

func (d *Driver) WhenITryToCreateTheRecipe(recipe domain.Recipe) {
	t := d.t
	t.Helper()

	err := d.Recipize.CreateRecipe(d.ctx, recipe)
	assert.NoError(t, err, "failed to create recipe")

	t.Cleanup(func() { assert.NoError(t, d.Recipize.DeleteRecipe(d.ctx, recipe.Name)) })
}

func (d *Driver) WhenIAddTheRecipeToMyMealPlan(name domain.RecipeName) {
	t := d.t
	t.Helper()

	err := d.Recipize.AddRecipeToMealPlan(d.ctx, name)
	assert.NoError(t, err, "failed to add recipe to meal plan")
}

func (d *Driver) ThenICanSeeTheRecipeInMyList(name domain.RecipeName) {
	t := d.t
	t.Helper()

	recipes, err := d.Recipize.GetRecipes(d.ctx)
	assert.NoError(t, err, "failed to get recipes")

	if !slices.Contains(recipes, name) {
		t.Fatalf("%q is not in the list of recipes: %v", name, recipes)
	}
}

func (d *Driver) ThenICanSeeTheIngredientsOfTheRecipe(name domain.RecipeName, expectedIngredients []domain.IngredientName) {
	t := d.t
	t.Helper()

	recipe, found, err := d.Recipize.GetRecipe(d.ctx, name)
	assert.NoError(t, err, "failed to get recipe")
	assert.True(t, found, "the recipe doesn't exist")

	var missingIngredients []domain.IngredientName
	for _, expectedIngredient := range expectedIngredients {
		if !slices.Contains(recipe.Ingredients, expectedIngredient) {
			missingIngredients = append(missingIngredients, expectedIngredient)
		}
	}

	if len(missingIngredients) > 0 {
		t.Fatalf("the ingredients %v are missing from the result %v", missingIngredients, recipe.Ingredients)
	}

	assert.Equal(t, len(expectedIngredients), len(recipe.Ingredients), "got the wrong amount of ingredients. expected %v but got %v", expectedIngredients, recipe.Ingredients)
}

func (d *Driver) ThenICanSeeTheRecipeInMyMealPlan(name domain.RecipeName) {
	t := d.t
	t.Helper()

	mealPlan, err := d.Recipize.GetMealPlan(d.ctx)
	assert.NoError(t, err, "failed to get recipes")

	if !slices.Contains(mealPlan.Recipes, name) {
		t.Fatalf("%q is not in the meal plan: %v", name, mealPlan.Recipes)
	}
}

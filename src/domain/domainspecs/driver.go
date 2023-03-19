package domainspecs

import (
	"context"
	"github.com/alecthomas/assert/v2"
	"github.com/tamj0rd2/recipize/src/domain"
	"github.com/tamj0rd2/recipize/src/ports"
	"github.com/tamj0rd2/recipize/testhelpers/slices"
	"testing"
)

func NewDriver(ctx context.Context, t testing.TB, recipize ports.Recipize) *Driver {
	return &Driver{
		ctx:      ctx,
		t:        t,
		Recipize: recipize,
	}
}

type Driver struct {
	t testing.TB
	ports.Recipize
	ctx context.Context
}

func (d *Driver) GivenTheRecipeDoesNotExist(name domain.RecipeName) {
	t := d.t
	t.Helper()

	exists, err := d.DoesRecipeExist(d.ctx, name)
	assert.NoError(t, err, "failed to check if recipe exists")
	assert.False(t, exists, "the recipe already exists")
}

func (d *Driver) WhenIAttemptToCreateTheRecipe(name domain.RecipeName) {
	t := d.t
	t.Helper()

	err := d.Recipize.CreateRecipe(d.ctx, name)
	assert.NoError(t, err, "failed to create recipe")
}

func (d *Driver) ThenICanSeeTheRecipeInMyRecipesList(name domain.RecipeName) {
	t := d.t
	t.Helper()

	d.thenTheRecipeExists(name)
	recipes, err := d.Recipize.GetRecipes(d.ctx)
	assert.NoError(t, err, "failed to get recipes")
	if !slices.Contains(recipes, name) {
		t.Fatalf("%q is not in the list of recipes: %v", name, recipes)
	}
}

func (d *Driver) thenTheRecipeExists(name domain.RecipeName) {
	t := d.t
	t.Helper()

	exists, err := d.DoesRecipeExist(d.ctx, name)
	assert.NoError(t, err, "failed to check if recipe exists")
	assert.True(t, exists, "the recipe doesn't exist")
}
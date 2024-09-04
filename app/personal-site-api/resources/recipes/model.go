package recipes

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// A struct to model the object
type Ingredient struct {
	Name        string   `json:"name"`
	Measurement *float64 `json:"measurement"`
	Units       *string  `json:"units"`
	Category    string   `json:"category"`
}

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Serves      *int         `json:"serves"`
	Url         *string      `json:"url"`
}

// An interface to refresent the Model (for mocking in test)
type RecipeDataAccessLayer interface {
	All() ([]Recipe, error)
}

// The Model with Database Implementation
type RecipeModel struct {
	DB *pgx.Conn
}

func (model *RecipeModel) All() (recipes []Recipe, err error) {
	type QueryResult struct {
		Ingredient
		Recipe
	}

	stmt := `
		SELECT r.name, r.serves, r.url, ri.ingredient, ri.measurement, ri.units, i.category
		FROM recipes r
		LEFT OUTER JOIN recipe_ingredients ri
		ON r.name = ri.recipe
		LEFT OUTER JOIN ingredients i
		on ri.ingredient = i.name
	`
	rows, err := model.DB.Query(context.Background(), stmt)
	if err != nil {
		return recipes, err
	}

	var recipe Recipe
	for rows.Next() {
		var res QueryResult

		err = rows.Scan(&res.Recipe.Name, &res.Serves, &res.Url, &res.Ingredient.Name, &res.Measurement, &res.Units, &res.Category)
		if err != nil {
			break
		}

		if res.Recipe.Name != recipe.Name {
			if recipe.Name != "" {
				// append if not nil recipe
				recipes = append(recipes, recipe)
			}
			// reset recipe each time it changes
			recipe = res.Recipe
		}

		recipe.Ingredients = append(recipe.Ingredients, res.Ingredient)
	}

	return recipes, err
}

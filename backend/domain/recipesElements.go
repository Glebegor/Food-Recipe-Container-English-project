package domain

type RecipesElement struct {
	Id       int    `json:"id"db:"id"`
	RecipeId int    `json:"recipeId"db:"recipe_id"`
	Name     string `json:"name"db:"name"`
	Quantity int    `json:"quantity"db:"quantity"`
	Unit     string `json:"unit"db:"unit"`
}

type IRecipesElementRepository interface {
}

type IRecipesElementService interface {
}

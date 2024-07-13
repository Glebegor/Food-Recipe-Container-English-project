package domain

type RecipesElement struct {
	Id       int    `json:"id"db:"id"`
	RecipeId int    `json:"recipeId"db:"recipe_id"`
	Name     string `json:"name"db:"name"`
	Quantity int    `json:"quantity"db:"quantity"`
	Unit     string `json:"unit"db:"unit"`
}

type IRecipesElementRepository interface {
	Get() (error, []RecipesElement)
	GetById(id int) (error, RecipesElement)
	Post(RecipesElement) error
	Delete(id int) error
}

type IRecipesElementService interface {
	Get() (error, []RecipesElement)
	GetById(id int) (error, RecipesElement)
	Post(RecipesElement) error
	Delete(id int) error
}

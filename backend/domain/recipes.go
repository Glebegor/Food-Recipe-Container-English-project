package domain

type Recipe struct {
	id          int    `json:"id"db:"id"`
	name        string `json:"name"db:"name"`
	description string `json:"description"db:"description"`
}

type IRecipeRepository interface {
}

type IRecipeService interface {
	Get() (error, []Recipe)
	GetById() (error, Recipe)
	Post() error
	Delete() error
}

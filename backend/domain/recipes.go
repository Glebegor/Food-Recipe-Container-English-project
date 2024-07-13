package domain

type Recipe struct {
	Id          int    `json:"id"db:"id"`
	Name        string `json:"name"db:"name"`
	Description string `json:"description"db:"description"`
}

type IRecipeRepository interface {
	Get() (error, []Recipe)
	GetById(id int) (error, Recipe)
	Post(Recipe) error
	Delete(id int) error
}

type IRecipeService interface {
	Get() (error, []Recipe)
	GetById(id int) (error, Recipe)
	Post(Recipe) error
	Delete(id int) error
}

package repositories

import "github.com/jmoiron/sqlx"

type RecipesRepository struct {
	Db *sqlx.DB
}

func (r *RecipesRepository) Get() {

}

func (r *RecipesRepository) GetById() {

}

func (r *RecipesRepository) Post() {

}

func (r *RecipesRepository) Delete() {

}

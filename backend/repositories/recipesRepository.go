package repositories

import "github.com/jmoiron/sqlx"

type RecipesRepository struct {
	Db *sqlx.DB
}

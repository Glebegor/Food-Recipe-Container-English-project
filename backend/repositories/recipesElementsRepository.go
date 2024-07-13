package repositories

import "github.com/jmoiron/sqlx"

type RecipesElementRepository struct {
	Db *sqlx.DB
}

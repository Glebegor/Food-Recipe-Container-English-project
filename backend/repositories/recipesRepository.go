package repositories

import (
	"backend/domain"
	"github.com/jmoiron/sqlx"
)

type RecipesRepository struct {
	Db *sqlx.DB
}

func (r *RecipesRepository) Get() (error, []domain.Recipe) {
	var data []domain.Recipe

	err := r.Db.Select(&data, "SELECT * FROM recipe")
	if err != nil {
		return err, nil
	}

	return nil, data
}

func (r *RecipesRepository) GetById(id int) (error, domain.Recipe) {
	return nil, domain.Recipe{}
}

func (r *RecipesRepository) Post(input domain.Recipe) error {
	_, err := r.Db.NamedExec("INSERT INTO recipe (name, description) VALUES (:name, :description)", input)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecipesRepository) Delete() error {
	return nil
}

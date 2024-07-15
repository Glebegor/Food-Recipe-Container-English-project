package repositories

import (
	"backend/domain"
	"errors"
	"github.com/jmoiron/sqlx"
)

type RecipesElementRepository struct {
	Db *sqlx.DB
}

func (r *RecipesElementRepository) Get() (error, []domain.RecipesElement) {
	var data []domain.RecipesElement

	err := r.Db.Select(&data, "SELECT * FROM recipe_element")
	if err != nil {
		return err, nil
	}

	return nil, data
}

func (r *RecipesElementRepository) GetById(id int) (error, domain.RecipesElement) {
	var data domain.RecipesElement

	err := r.Db.Get(&data, "SELECT * FROM recipe_element WHERE id = $1", id)
	if err != nil {
		return err, domain.RecipesElement{}
	}
	return nil, data
}

func (r *RecipesElementRepository) Post(input domain.RecipesElement) error {
	// Check if the recipe_id exists in the recipes table
	var count int
	err := r.Db.Get(&count, "SELECT COUNT(*) FROM recipe WHERE id = $1", input.RecipeId)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("recipe_id does not exist in recipes table")
	}

	// Insert into recipe_element table
	_, err = r.Db.NamedExec("INSERT INTO recipe_element (recipe_id, quantity, name, unit) VALUES (:recipe_id, :quantity, :name, :unit)", input)
	if err != nil {
		return err
	}

	return nil
}

func (r *RecipesElementRepository) Delete(id int) error {
	_, err := r.Db.Exec("DELETE FROM recipe_element WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

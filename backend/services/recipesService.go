package services

import "backend/domain"

type RecipesService struct {
	Repo domain.IRecipeRepository
}

func (rs *RecipesService) Get() (error, []domain.Recipe) {
	return rs.Repo.Get()
}

func (rs *RecipesService) GetById(id int) (error, domain.Recipe) {
	return rs.Repo.GetById(id)
}

func (rs *RecipesService) Post(input domain.Recipe) error {
	return rs.Repo.Post(input)
}

func (rs *RecipesService) Delete() error {
	return nil
}

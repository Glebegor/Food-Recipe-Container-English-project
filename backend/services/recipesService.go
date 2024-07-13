package services

import "backend/domain"

type RecipesService struct {
	Repo domain.IRecipeRepository
}

func (rs *RecipesService) Get() (error, []domain.Recipe) {
	return nil, nil
}

func (rs *RecipesService) GetById() (error, domain.Recipe) {
	return nil, nil
}

func (rs *RecipesService) Post() error {
	return nil
}

func (rs *RecipesService) Delete() error {
	return nil
}

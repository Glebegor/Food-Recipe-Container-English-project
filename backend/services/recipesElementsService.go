package services

import "backend/domain"

type RecipesElementService struct {
	Repo domain.IRecipesElementRepository
}

func (rs *RecipesElementService) Get() (error, []domain.RecipesElement) {
	return rs.Repo.Get()
}

func (rs *RecipesElementService) GetById(id int) (error, domain.RecipesElement) {
	return rs.Repo.GetById(id)
}

func (rs *RecipesElementService) Post(input domain.RecipesElement) error {
	return rs.Repo.Post(input)
}

func (rs *RecipesElementService) Delete(id int) error {
	return rs.Repo.Delete(id)
}

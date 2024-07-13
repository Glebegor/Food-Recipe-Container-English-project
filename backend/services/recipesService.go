package services

import "backend/domain"

type RecipesService struct {
	Repo domain.IRecipeRepository
}

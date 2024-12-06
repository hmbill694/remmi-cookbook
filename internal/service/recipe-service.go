package service

import (
	"remmi-cookbook/internal/database"
	"remmi-cookbook/internal/repository"
)

type RecipeService struct {
	DbInstance database.DBInstance
}

func NewRecipeService(dbInstance database.DBInstance) RecipeService {
	return RecipeService{
		DbInstance: dbInstance,
	}
}

func (rs RecipeService) GetRecipesForUser(search string) ([]repository.Recipe, error) {
	repo := repository.New(rs.DbInstance.Db)

	if search == "" {
		return repo.GetAllRecipes(*rs.DbInstance.DbContext)
	}

	return repo.SearchForRecipes(*rs.DbInstance.DbContext, search)
}

// func (rs RecipeService) GetRecipesForUserHtml() (string, error) {
// 	recipes, err := rs.GetRecipesForUser()

// 	if err != nil {
// 		return "", err
// 	}

// }

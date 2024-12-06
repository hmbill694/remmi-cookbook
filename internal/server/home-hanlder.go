package server

import (
	"log/slog"
	homeView "remmi-cookbook/internal/views/home"

	"github.com/labstack/echo/v4"
)

func (s *Server) HomeHandler(e echo.Context) error {
	searchParam := e.QueryParam("search")

	recipes, err := s.recipeService.GetRecipesForUser(searchParam)

	home := homeView.HomePage(homeView.HomePageProps{
		Recipes: recipes,
		UserId:  "foo",
		Err:     err,
	})

	return home.Render(e.Request().Context(), e.Response().Writer)
}

func (s *Server) HomeHandlerSearch(e echo.Context) error {
	searchParam := e.QueryParam("search")

	recipes, err := s.recipeService.GetRecipesForUser(searchParam)

	if err != nil {
		slog.Info(err.Error())
		return err
	}

	cards := homeView.RecipeCardList(homeView.RecipeCardListProps{
		Recipes: recipes,
		Err:     err,
	})

	return cards.Render(e.Request().Context(), e.Response().Writer)
}

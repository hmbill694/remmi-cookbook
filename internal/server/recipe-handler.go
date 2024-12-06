package server

import (
	"log/slog"
	rp "remmi-cookbook/internal/views/recipe-page"

	"github.com/labstack/echo/v4"
)

func (s *Server) RecipePageHandler(e echo.Context) error {

	recipeId := e.Param("id")
	slog.Info(recipeId)
	slog.Info("FUCK")

	// recipeDetails, err := s.recipeService.GetRecipesForUser(recipeId)

	// if err != nil {
	// 	slog.Info(err.Error())
	// 	return err
	// }

	recipePage := rp.RecipePage()

	return TempleToEchoResponse(recipePage, e)

}

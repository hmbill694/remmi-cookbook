package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"remmi-cookbook/internal/database"
	"remmi-cookbook/internal/service"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port          int
	recipeService service.RecipeService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	ctx := context.Background()

	dbInstance := database.New(ctx)

	recipeService := service.NewRecipeService(*dbInstance)

	NewServer := &Server{
		port:          port,
		recipeService: recipeService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

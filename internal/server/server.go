package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"remmi-cookbook/internal/database"
)

type Server struct {
	port int

	Db database.Service
}

func NewServer() (*http.Server, database.Service) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db := database.New()

	NewServer := &Server{
		port: port,
		Db:   db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, db
}

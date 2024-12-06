package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"remmi-cookbook/cmd/web"

	notFoundPage "remmi-cookbook/internal/views/not-found-page"

	"github.com/coder/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func errorHandler(err error, e echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code == 404 {
		if renderError := TempleToEchoResponse(notFoundPage.NotFoundPage(), e); renderError != nil {
			slog.Error(fmt.Sprintf("Failed to render 404 response %s", renderError.Error()))
		}
	}

}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/", s.HomeHandler)

	e.GET("/api/recipe/search", s.HomeHandlerSearch)

	e.GET("/recipe/:id", s.RecipePageHandler)

	e.GET("/health/up", s.up)

	e.GET("/websocket", s.websocketHandler)

	e.HTTPErrorHandler = errorHandler

	return e
}

func (s *Server) up(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (s *Server) websocketHandler(c echo.Context) error {
	w := c.Response().Writer
	r := c.Request()
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

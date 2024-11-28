package server

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func HanlderFromTempl(comp templ.Component) echo.HandlerFunc {
 return echo.WrapHandler(templ.Handler(comp))
}

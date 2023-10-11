package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryuji-cre8ive/super-suica/internal/usecase"
)

type Handler struct {
	UserHandler
}

func New(u *usecase.Usecase) *Handler {
	return &Handler{
		UserHandler: &userHandler{u.User},
	}
}

func SetApi(e *echo.Echo, h *Handler) {
	g := e.Group("/api/v1")
	g.GET("/healthcheck", HealthCheckHandler)
	g.POST("/webhook", h.UserHandler.Create)

}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	return e
}

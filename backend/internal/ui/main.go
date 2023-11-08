package ui

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryuji-cre8ive/super-metro/internal/usecase"
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
	g.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	g.GET("/healthcheck", HealthCheckHandler)
	g.POST("/signup", h.UserHandler.Create)
	g.POST("/login", h.UserHandler.Login)
	g.POST("/logout", h.UserHandler.Logout)
	g.POST("/top-up", h.UserHandler.TopUp)
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

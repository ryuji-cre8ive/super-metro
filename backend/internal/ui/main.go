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
	PaymentHandler
	TransactionHandler
}

func New(u *usecase.Usecase) *Handler {
	return &Handler{
		UserHandler:        &userHandler{u.User, u.Payment, u.Transaction},
		PaymentHandler:     &paymentHandler{u.Payment},
		TransactionHandler: &transactionHandler{u.Transaction},
	}
}

func SetApi(e *echo.Echo, h *Handler) {
	g := e.Group("/api/v1")
	authGroup := e.Group("/api/v1")
	authGroup.Use(h.UserHandler.CheckCookieExpiration())
	g.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	g.GET("/healthcheck", HealthCheckHandler)
	g.POST("/signup", h.UserHandler.Create)
	g.POST("/login", h.UserHandler.Login)
	g.GET("/cookie/:userID", h.UserHandler.Cookie)
	authGroup.POST("/logout", h.UserHandler.Logout)
	authGroup.POST("/top-up", h.UserHandler.TopUp)

	authGroup.GET("/credit-card/:userID", h.PaymentHandler.Get)
	authGroup.POST("/credit-card/add", h.PaymentHandler.Add)

	authGroup.GET("/amount/:userID", h.UserHandler.GetAmount)

	authGroup.GET("/transaction/:userID", h.TransactionHandler.Get)
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

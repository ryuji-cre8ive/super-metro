package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-suica/internal/usecase"

	"golang.org/x/xerrors"
)

type (
	UserHandler interface {
		Create(c echo.Context) error
	}

	userHandler struct {
		usecase.UserUsecase
	}
)

func (h *userHandler) Create(c echo.Context) error {
	userName := c.Param("userName")
	password := c.Param("password")
	err := h.UserUsecase.Create(c, userName, password)
	if err != nil {
		return xerrors.Errorf("failed to post User: %w", err)
	}
	return c.NoContent(200)
}

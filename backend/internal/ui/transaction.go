package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/usecase"
	"golang.org/x/xerrors"
)

type (
	TransactionHandler interface {
		Get(c echo.Context) error
	}

	transactionHandler struct {
		usecase.TransactionUsecase
	}
)

func (h *transactionHandler) Get(c echo.Context) error {
	userId := c.Param("userID")
	transactions, transactionErr := h.TransactionUsecase.Get(c, userId)
	if transactionErr != nil {
		return xerrors.Errorf("failed to get Transaction: %w", transactionErr)
	}

	return c.JSON(200, transactions)
}

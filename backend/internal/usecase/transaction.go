package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/stores"
)

type (
	TransactionUsecase interface {
		Add(c echo.Context, userId string, paymentId string, transactionType string, amount int) error
		Get(c echo.Context, userId string) ([]*domain.Transaction, error)
	}

	transactionUsecase struct {
		stores *stores.Stores
	}
)

func (u *transactionUsecase) Add(c echo.Context, userId string, paymentId string, transactionType string, amount int) error {
	return u.stores.Transaction.Add(userId, paymentId, transactionType, amount)
}

func (u *transactionUsecase) Get(c echo.Context, userId string) ([]*domain.Transaction, error) {
	return u.stores.Transaction.Get(userId)
}

// func (u *transactionUsecase) Delete(c echo.Context, userId string) error {
// 	return u.stores.Transaction.Delete(userId)
// }

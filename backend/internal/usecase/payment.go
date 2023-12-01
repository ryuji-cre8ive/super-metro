package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/stores"
)

type (
	PaymentUsecase interface {
		Add(c echo.Context, userId string, cardNumber string, expiryDate string, cvv string) error
		Delete(c echo.Context, userId string) error
		Get(c echo.Context, userId string) (*domain.Payment, error)
	}

	paymentUsecase struct {
		stores *stores.Stores
	}
)

func (u *paymentUsecase) Add(c echo.Context, userId string, cardNumber string, expiryDate string, cvv string) error {
	return u.stores.Payment.Add(userId, cardNumber, expiryDate, cvv)
}

func (u *paymentUsecase) Delete(c echo.Context, userId string) error {
	return u.stores.Payment.Delete(userId)
}

func (u *paymentUsecase) Get(c echo.Context, userId string) (*domain.Payment, error) {
	return u.stores.Payment.Get(userId)
}

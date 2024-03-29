package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/stores"
	"github.com/ryuji-cre8ive/super-metro/internal/utils"
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
	key := []byte(userId)[:32] // 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256
	encryptedCardNumber, err := utils.Encrypt([]byte(cardNumber), key)
	if err != nil {
		return err
	}

	encryptedExpiryDate, err := utils.Encrypt([]byte(expiryDate), key)
	if err != nil {
		return err
	}

	encryptedCVV, err := utils.Encrypt([]byte(cvv), key)
	if err != nil {
		return err
	}
	return u.stores.Payment.Add(userId, string(encryptedCardNumber), string(encryptedExpiryDate), string(encryptedCVV))
}

func (u *paymentUsecase) Delete(c echo.Context, userId string) error {
	return u.stores.Payment.Delete(userId)
}

func (u *paymentUsecase) Get(c echo.Context, userId string) (*domain.Payment, error) {
	userPayment, err := u.stores.Payment.Get(userId)
	if userPayment.CardNumber == "" {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	key := []byte(userId)[:32] // AES-256を選択するために32バイトを使用します。userIDはuuidで、36バイトですが、最初の32バイトだけを使用します。

	decryptedCardNumber, err := utils.Decrypt(userPayment.CardNumber, key)
	if err != nil {
		return nil, err
	}

	decryptedExpiryDate, err := utils.Decrypt(userPayment.ExpiryDate, key)
	if err != nil {
		return nil, err
	}

	decryptedCVV, err := utils.Decrypt(userPayment.CVV, key)
	if err != nil {
		return nil, err
	}

	userPayment.CardNumber = string(decryptedCardNumber)
	userPayment.ExpiryDate = string(decryptedExpiryDate)
	userPayment.CVV = string(decryptedCVV)

	return userPayment, nil
}

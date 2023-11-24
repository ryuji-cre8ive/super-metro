package stores

import (
	"time"

	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"gorm.io/gorm"
)
type (
	PaymentStore interface {
		Add(userId string, cardNumber string, expiryDate string, cvv string) error
	}

	paymentStore struct {
		*gorm.DB
	}
)

func(s *paymentStore) Add(userId string, cardNumber string, expiryDate string, cvv string) error {
	return s.DB.Create(&domain.Payment{
		UserID: userId,
		CardNumber: cardNumber,
		ExpiryDate: expiryDate,
		CVV: cvv,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}
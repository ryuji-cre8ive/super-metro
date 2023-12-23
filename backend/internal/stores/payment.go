package stores

import (
	"time"

	"github.com/google/uuid"

	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"gorm.io/gorm"
)

type (
	PaymentStore interface {
		Add(userId string, cardNumber string, expiryDate string, cvv string) error
		Delete(userId string) error
		Get(userId string) (*domain.Payment, error)
	}

	paymentStore struct {
		*gorm.DB
	}
)

func (s *paymentStore) Add(userId string, cardNumber string, expiryDate string, cvv string) error {
	return s.DB.Create(&domain.Payment{
		ID:         uuid.Must(uuid.NewRandom()).String(),
		UserID:     userId,
		CardNumber: cardNumber,
		ExpiryDate: expiryDate,
		CVV:        cvv,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}).Error
}

func (s *paymentStore) Delete(userId string) error {
	return s.DB.Model(&domain.Payment{}).Where("user_id = ?", userId).Update("DeletedAt", time.Now()).Error
}

func (s *paymentStore) Get(userId string) (*domain.Payment, error) {
	var payment *domain.Payment
	result := s.DB.Where("user_id = ? AND deleted_at IS NULL", userId).First(&payment)
	if result.Error != nil {
		return nil, result.Error
	}
	return payment, nil
}

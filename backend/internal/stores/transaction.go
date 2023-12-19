package stores

import (
	"time"

	"github.com/google/uuid"

	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"gorm.io/gorm"
)

type (
	TransactionStore interface {
		Add(userId string, paymentId string, transactionType string, amount int) error
		Get(userId string) ([]*domain.Transaction, error)
	}

	transactionStore struct {
		*gorm.DB
	}
)

func (s *transactionStore) Add(userId string, paymentId string, transactionType string, amount int) error {
	return s.DB.Create(&domain.Transaction{
		ID:              uuid.Must(uuid.NewRandom()).String(),
		UserID:          userId,
		PaymentID:       paymentId,
		Amount:          amount,
		TransactionType: transactionType,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}).Error
}

func (s *transactionStore) Get(userId string) ([]*domain.Transaction, error) {
	var transactions []*domain.Transaction
	err := s.DB.Where("user_id = ?", userId).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

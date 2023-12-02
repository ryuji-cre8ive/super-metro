package stores

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/utils"
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
	result := s.DB.Where("user_id = ? AND deleted_at IS NULL", userId).Find(&payment)
	fmt.Println("result", payment)
	if payment.ID == "" {
		return nil, nil
	}
	key := []byte(userId)[:32] // AES-256を選択するために32バイトを使用します。userIDはuuidで、36バイトですが、最初の32バイトだけを使用します。

	decryptedCardNumber, err := utils.Decrypt(payment.CardNumber, key)
	if err != nil {
		return nil, err
	}

	decryptedExpiryDate, err := utils.Decrypt(payment.ExpiryDate, key)
	if err != nil {
		return nil, err
	}

	decryptedCVV, err := utils.Decrypt(payment.CVV, key)
	if err != nil {
		return nil, err
	}

	payment.CardNumber = string(decryptedCardNumber)
	payment.ExpiryDate = string(decryptedExpiryDate)
	payment.CVV = string(decryptedCVV)
	return payment, result.Error
}

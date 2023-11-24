package stores

type (
	PaymentStore interface {
		Add(c echo.Context, cardNumber string, expiryDate string, cvv string) error
	}

	paymentStore struct {
		*gorm.DB
	}
)

func(s *paymentStore) Add(c echo.Context, cardNumber string, expiryDate string, cvv string) error {
	return s.DB.Create(&domain.Payment{
		CardNumber: cardNumber,
		ExpiryDate: expiryDate,
		CVV: cvv,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}
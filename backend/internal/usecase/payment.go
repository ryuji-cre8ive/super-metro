package usecase

type (
	PaymentUsecase interface {
		Add(c echo.Context) error
	}

	paymentUsecase struct {
		stores *stores.Stores
	}
)

func (u *paymentUsecase) Add(c echo.Context, cardNumber string, expiryDate string, cvv string) error {
	return u.stores.Payment.Add(cardNumber, expiryDate, cvv)
}
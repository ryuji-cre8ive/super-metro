import ()


type (
	PaymentHandler interface {
		Add(c echo.Context) error
	}

	userHandler struct {
		usecase.PaymentUsecase
	}
)

func (h *paymentHandler) Add(c echo.Context) error {
	param := new(domain.Payment)
	if err := c.Bind(param); err != nil {
		return xerrors.Errorf("failed to bind Payment: %w", err)
	}

	cardNumber := param.CardNumber
	if cardNumber == "" {
		return xerrors.Errorf("cardNumber is empty")
	}
	if len(cardNumber) != 16 {
		return xerrors.Errorf("cardNumber is invalid")
	}

	expiryDate := param.ExpiryDate
	if expiryDate == "" {
		return xerrors.Errorf("expiryDate is empty")
	}
	if len(expiryDate) != 4 {
		return xerrors.Errorf("expiryDate is invalid")
	}
	
	cvv := param.CVV
	if cvv == "" {
		return xerrors.Errorf("cvv is empty")
	}
	if len(cvv) != 3 {
		return xerrors.Errorf("cvv is invalid")
	}

	userId := param.UserId
	paymentErr := h.PaymentUsecase.Add(c, userId, cardNumber, expiryDate, cvv)
	if paymentErr != nil {
		return xerrors.Errorf("failed to post Payment: %w", paymentErr)
	}

	return c.String(200, "success")
}
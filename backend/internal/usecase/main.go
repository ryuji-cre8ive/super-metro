package usecase

import "github.com/ryuji-cre8ive/super-metro/internal/stores"

type Usecase struct {
	User        UserUsecase
	Payment     PaymentUsecase
	Transaction TransactionUsecase
}

func New(s *stores.Stores) *Usecase {
	return &Usecase{
		User:        &userUsecase{stores: s},
		Payment:     &paymentUsecase{stores: s},
		Transaction: &transactionUsecase{stores: s},
	}
}

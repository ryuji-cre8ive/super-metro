package stores

import (
	"gorm.io/gorm"
)

type Stores struct {
	DB          *gorm.DB
	User        UserStore
	Payment     PaymentStore
	Transaction TransactionStore
}

func New(db *gorm.DB) *Stores {
	return &Stores{
		DB:          db,
		User:        &userStore{db},
		Payment:     &paymentStore{db},
		Transaction: &transactionStore{db},
	}
}

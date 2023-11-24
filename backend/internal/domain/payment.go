package domain

import (
	"time"
)

type Payment struct {
	ID         string    `db:"id"`
	UserID     string    `db:"userId"`
	CardNumber string    `db:"cardNumber"`
	ExpiryDate string    `db:"expiryDate"`
	CVV        string    `db:"cvv"`
	CreatedAt  time.Time `db:"createdAt"`
	UpdatedAt  time.Time `db:"updatedAt"`
	DeletedAt  *time.Time `db:"deletedAt"`
}
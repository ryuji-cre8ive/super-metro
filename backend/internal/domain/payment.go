package domain

import (
	"time"
)

type Transaction struct {
	ID              string     `db:"id" json:"id"`
	UserID          string     `db:"userId" json:"userId"`
	PaymentID       string     `db:"paymentId" json:"paymentId"`
	Amount          int        `db:"amount" json:"amount"`
	TransactionType string     `db:"transactionType" json:"transactionType"`
	CreatedAt       time.Time  `db:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time  `db:"updatedAt" json:"updatedAt"`
	DeletedAt       *time.Time `db:"deletedAt" json:"deletedAt"`
}

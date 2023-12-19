package stores

import (
	"errors"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPaymentStore_Add(t *testing.T) {
	type input struct {
		userId     string
		cardNumber string
		expiryDate string
		cvv        string
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				userId:     "1",
				cardNumber: "1234567812345678",
				expiryDate: "12/24",
				cvv:        "123",
			},
			wantErr: false,
		},
		"failed": {
			input: input{
				userId:     "2",
				cardNumber: "1234567812345678",
				expiryDate: "12/24",
				cvv:        "123",
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sqlDB, mock, _ := sqlmock.New()
			defer sqlDB.Close()

			db, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

			s := &paymentStore{
				DB: db,
			}

			mock.ExpectBegin()
			if tt.wantErr {
				mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			} else {
				mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			err := s.Add(tt.input.userId, tt.input.cardNumber, tt.input.expiryDate, tt.input.cvv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

package stores

import (
	"errors"
	"testing"
	"time"

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

func TestPaymentStore_Delete(t *testing.T) {
	type input struct {
		userId string
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				userId: "1",
			},
			wantErr: false,
		},
		"failed": {
			input: input{
				userId: "2",
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

			mock.ExpectBegin() // これを追加します
			if tt.wantErr {
				mock.ExpectExec(".*UPDATE \"payments\" SET \"deleted_at\"=\\$1,\"updated_at\"=\\$2 WHERE user_id = \\$3.*").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), tt.input.userId).
					WillReturnError(errors.New("database error"))
				mock.ExpectRollback()
			} else {
				mock.ExpectExec(".*UPDATE \"payments\" SET \"deleted_at\"=\\$1,\"updated_at\"=\\$2 WHERE user_id = \\$3.*").
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), tt.input.userId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			err := s.Delete(tt.input.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPaymentStore_Get(t *testing.T) {
	type input struct {
		userId string
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				userId: "1",
			},
			wantErr: false,
		},
		"failed": {
			input: input{
				userId: "2",
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

			if tt.wantErr {
				mock.ExpectQuery(".*SELECT \\* FROM \"payments\" WHERE user_id = \\$1 AND deleted_at IS NULL ORDER BY \"payments\".\"id\" LIMIT 1.*").
					WithArgs(tt.input.userId).
					WillReturnError(errors.New("database error"))
			} else {
				rows := sqlmock.NewRows([]string{"id", "user_id", "card_number", "expiry_date", "cvv", "created_at", "updated_at", "deleted_at"}).
					AddRow("1", "1", "encryptedCardNumber", "encryptedExpiryDate", "encryptedCVV", time.Now(), time.Now(), nil)
				mock.ExpectQuery("SELECT \\* FROM \"payments\" WHERE user_id = \\$1 AND deleted_at IS NULL ORDER BY \"payments\".\"id\" LIMIT 1").
					WithArgs(tt.input.userId).
					WillReturnRows(rows)
			}
			_, err := s.Get(tt.input.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

package stores

import (
	"testing"

	"golang.org/x/xerrors"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestTransactionStore_Add(t *testing.T) {
	type input struct {
		userId          string
		paymentId       string
		transactionType string
		amount          int
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				userId:          "1",
				paymentId:       "1",
				transactionType: "debit",
				amount:          100,
			},
			wantErr: false,
		},
		"failed": {
			input: input{
				userId:          "2",
				paymentId:       "2",
				transactionType: "credit",
				amount:          200,
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sqlDB, mock, _ := sqlmock.New()
			defer sqlDB.Close()

			db, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

			s := &transactionStore{
				DB: db,
			}

			if tt.wantErr {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO").WillReturnError(xerrors.New("database error"))
				mock.ExpectRollback()
			} else {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			}

			err := s.Add(tt.input.userId, tt.input.paymentId, tt.input.transactionType, tt.input.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

// func TestTransactionStore_Get(t *testing.T) {
// 	type input struct {
// 		userId string
// 	}

// 	tests := map[string]struct {
// 		input   input
// 		wantErr bool
// 	}{
// 		"success": {
// 			input: input{
// 				userId: "1",
// 			},
// 			wantErr: false,
// 		},
// 		"failed": {
// 			input: input{
// 				userId: "2",
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			sqlDB, mock, _ := sqlmock.New()
// 			defer sqlDB.Close()

// 			db, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

// 			s := &transactionStore{
// 				DB: db,
// 			}

// 			if tt.wantErr {
// 				mock.ExpectQuery("^SELECT (.+) FROM \"transactions\" WHERE user_id = (.+)").
// 					WithArgs(tt.input.userId).
// 					WillReturnError(errors.New("database error"))
// 			} else {
// 				rows := sqlmock.NewRows([]string{"id", "user_id", "payment_id", "amount", "transaction_type", "created_at", "updated_at"}).
// 					AddRow(uuid.Must(uuid.NewRandom()).String(), tt.input.userId, "1", 100, "debit", time.Now(), time.Now())
// 				mock.ExpectQuery("^SELECT (.+) FROM \"transactions\" WHERE user_id = (.+)").
// 					WithArgs(tt.input.userId).
// 					WillReturnRows(rows)
// 			}

// 			_, err := s.Get(tt.input.userId)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

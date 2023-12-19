package stores

import (
	"errors"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserStore_Create(t *testing.T) {
	type input struct {
		email    string
		userName string
		password string
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				email:    "test@test.com",
				userName: "test",
				password: "test",
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sqlDB, mock, _ := sqlmock.New()
			defer sqlDB.Close()

			db, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

			s := &userStore{
				DB: db,
			}

			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			err := s.Create(tt.input.email, tt.input.userName, tt.input.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserStore_FindByEmail(t *testing.T) {
	type input struct {
		email string
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				email: "test@test.com",
			},
			wantErr: false,
		},
		"failed": {
			input: input{
				email: "test@test.com",
			},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			sqlDB, mock, _ := sqlmock.New()
			defer sqlDB.Close()

			db, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

			s := &userStore{
				DB: db,
			}

			if tt.wantErr {
				mock.ExpectQuery("^SELECT (.+) FROM \"users\" WHERE email = (.+) ORDER BY \"users\".\"id\" LIMIT 1").WithArgs(tt.input.email).WillReturnError(errors.New("database error"))
			} else {
				rows := sqlmock.NewRows([]string{"email"}).AddRow(tt.input.email)
				mock.ExpectQuery("^SELECT (.+) FROM \"users\" WHERE email = (.+) ORDER BY \"users\".\"id\" LIMIT 1").WithArgs(tt.input.email).WillReturnRows(rows)
			}

			user, err := s.FindByEmail(tt.input.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil && user.Email != tt.input.email {
				t.Errorf("FindByEmail() = %v, want %v", user.Email, tt.input.email)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

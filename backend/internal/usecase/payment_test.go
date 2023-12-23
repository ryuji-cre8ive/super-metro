package usecase

import (
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/stores"
	mock "github.com/ryuji-cre8ive/super-metro/internal/stores/mock"
	utilMock "github.com/ryuji-cre8ive/super-metro/internal/utils/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
)

func TestPaymentUsecase_Add(t *testing.T) {
	type input struct {
		ctx        echo.Context
		userId     string
		cardNumber string
		expiryDate string
		cvv        string
	}

	// wantedError := xerrors.New("error")

	tests := map[string]struct {
		input    input
		want     error
		mockFunc func(m *mock.MockPaymentStore, u *utilMock.MockEncryptType)
	}{
		"success": {
			input: input{
				ctx:        nil,
				userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				cardNumber: "1234567890123456",
				expiryDate: "1234",
				cvv:        "123",
			},
			want: nil,
			mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncryptType) {
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		// "failed: standard-error": {
		// 	input: input{
		// 		ctx:        nil,
		// 		userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
		// 		cardNumber: "1234567890123456",
		// 		expiryDate: "1234",
		// 		cvv:        "123",
		// 	},
		// 	want: wantedError,
		// 	mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncryptType) {
		// 		u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", nil).Times(3)
		// 		m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(wantedError)
		// 	},
		// },
		// "failed: cardNumber-encryption": {
		// 	input: input{
		// 		ctx:        nil,
		// 		userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
		// 		cardNumber: "1234567890123456",
		// 		expiryDate: "1234",
		// 		cvv:        "123",
		// 	},
		// 	want: wantedError,
		// 	mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncryptType) {
		// 		u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", wantedError).Times(1)
		// 		m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
		// 	},
		// },
		// "failed: expiryDate-encryption": {
		// 	input: input{
		// 		ctx:        nil,
		// 		userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
		// 		cardNumber: "1234567890123456",
		// 		expiryDate: "1234",
		// 		cvv:        "123",
		// 	},
		// 	want: wantedError,
		// 	mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncryptType) {
		// 		u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", nil).Times(1)
		// 		u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", wantedError).Times(1)
		// 		m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
		// 	},
		// },
		// "failed: cvv-encryption": {
		// 	input: input{
		// 		ctx:        nil,
		// 		userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
		// 		cardNumber: "1234567890123456",
		// 		expiryDate: "1234",
		// 		cvv:        "123",
		// 	},
		// 	want: wantedError,
		// 	mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncryptType) {
		// 		u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", nil).Times(2)
		// 		u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", wantedError).Times(1)
		// 		m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
		// 	},
		// },
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockPaymentStore(ctrl)
			encrypt := utilMock.NewMockEncryptType(ctrl)
			tt.mockFunc(store, encrypt)

			u := &paymentUsecase{
				stores: &stores.Stores{
					Payment: store,
				},
			}

			err := u.Add(tt.input.ctx, tt.input.userId, tt.input.cardNumber, tt.input.expiryDate, tt.input.cvv)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestPaymentUsecase_Delete(t *testing.T) {
	type input struct {
		ctx    echo.Context
		userId string
	}

	wantedError := xerrors.New("error")

	tests := map[string]struct {
		input    input
		want     error
		mockFunc func(m *mock.MockPaymentStore)
	}{
		"success": {
			input: input{
				ctx:    nil,
				userId: "6d1c3e1b-d068-431f-b188-a436ac54ce52",
			},
			want: nil,
			mockFunc: func(m *mock.MockPaymentStore) {
				m.EXPECT().Delete(gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: standard-error": {
			input: input{
				ctx:    nil,
				userId: "6d1c3e1b-d068-431f-b188-a436ac54ce52",
			},
			want: wantedError,
			mockFunc: func(m *mock.MockPaymentStore) {
				m.EXPECT().Delete(gomock.Any()).Return(wantedError)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockPaymentStore(ctrl)
			tt.mockFunc(store)

			u := &paymentUsecase{
				stores: &stores.Stores{
					Payment: store,
				},
			}

			err := u.Delete(tt.input.ctx, tt.input.userId)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestPaymentUsecase_Get(t *testing.T) {
	type input struct {
		ctx    echo.Context
		userId string
	}

	tests := map[string]struct {
		input    input
		want     *domain.Payment
		wantErr  error
		mockFunc func(m *mock.MockPaymentStore)
	}{
		"success": {
			input: input{
				ctx:    nil,
				userId: "6d1c3e1b-d068-431f-b188-a436ac54ce52",
			},
			want: &domain.Payment{
				UserID:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				CardNumber: "2222333344445555",
				ExpiryDate: "1122",
				CVV:        "222",
				CreatedAt:  time.Date(2023, time.December, 2, 23, 3, 51, 115734000, time.Local),
				UpdatedAt:  time.Date(2023, time.December, 2, 23, 3, 51, 115734000, time.Local),
			},
			wantErr: nil,
			mockFunc: func(m *mock.MockPaymentStore) {
				m.EXPECT().Get("6d1c3e1b-d068-431f-b188-a436ac54ce52").Return(&domain.Payment{
					UserID:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
					CardNumber: "oijncKLVhcNUbunWUxxr6adS7B3O0IYK9DkDIa8xtxI=",
					ExpiryDate: "w2Q4Cz8IzR6haGUHfDKceTgPysk=",
					CVV:        "A2bHOk9KVG3AnPhT3NRijNOgFQ==",
					CreatedAt:  time.Date(2023, time.December, 2, 23, 3, 51, 115734000, time.Local),
					UpdatedAt:  time.Date(2023, time.December, 2, 23, 3, 51, 115734000, time.Local),
				}, nil)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockPaymentStore(ctrl)
			tt.mockFunc(store)

			u := &paymentUsecase{
				stores: &stores.Stores{
					Payment: store,
				},
			}

			got, err := u.Get(tt.input.ctx, tt.input.userId)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.want, got)
			}

		})
	}
}

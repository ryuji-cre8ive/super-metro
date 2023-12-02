package usecase

import (
	"testing"

	"github.com/labstack/echo/v4"
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

	wantedError := xerrors.New("error")

	tests := map[string]struct {
		input    input
		want     error
		mockFunc func(m *mock.MockPaymentStore, u *utilMock.MockEncrypt)
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
			mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncrypt) {
				u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", nil).Times(3)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: standard-error": {
			input: input{
				ctx:        nil,
				userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				cardNumber: "1234567890123456",
				expiryDate: "1234",
				cvv:        "123",
			},
			want: wantedError,
			mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncrypt) {
				u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", nil).Times(3)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(wantedError)
			},
		},
		"failed: cardNumber-encryption": {
			input: input{
				ctx:        nil,
				userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				cardNumber: "1234567890123456",
				expiryDate: "1234",
				cvv:        "123",
			},
			want: wantedError,
			mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncrypt) {
				u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", wantedError).Times(1)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
		"failed: expiryDate-encryption": {
			input: input{
				ctx:        nil,
				userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				cardNumber: "1234567890123456",
				expiryDate: "1234",
				cvv:        "123",
			},
			want: wantedError,
			mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncrypt) {
				u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", nil).Times(1)
				u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", wantedError).Times(1)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
		"failed: cvv-encryption": {
			input: input{
				ctx:        nil,
				userId:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				cardNumber: "1234567890123456",
				expiryDate: "1234",
				cvv:        "123",
			},
			want: wantedError,
			mockFunc: func(m *mock.MockPaymentStore, u *utilMock.MockEncrypt) {
				u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", nil).Times(2)
				u.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return("", wantedError).Times(1)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockPaymentStore(ctrl)
			encrypt := utilMock.NewMockEncrypt(ctrl)
			tt.mockFunc(store, encrypt)

			u := &paymentUsecase{
				stores: &stores.Stores{
					Payment: store,
				},
				encrypt: encrypt,
			}

			err := u.Add(tt.input.ctx, tt.input.userId, tt.input.cardNumber, tt.input.expiryDate, tt.input.cvv)
			assert.Equal(t, tt.want, err)
		})
	}
}

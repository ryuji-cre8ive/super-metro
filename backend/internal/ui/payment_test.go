package ui

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	mock "github.com/ryuji-cre8ive/super-metro/internal/usecase/mock"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
)

func TestPaymentHandler_Add(t *testing.T) {
	type input struct {
		ctx        *echo.Echo
		cardNumber string
		expiryDate string
		cvv        string
		userId     string
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockPaymentUsecase)
	}{
		"success": {
			input: input{
				ctx:        echo.New(),
				cardNumber: "1234567812345678",
				expiryDate: "1122",
				cvv:        "123",
				userId:     "testID",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockPaymentUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"success with delete": {
			input: input{
				ctx:        echo.New(),
				cardNumber: "1234567812345678",
				expiryDate: "1122",
				cvv:        "123",
				userId:     "testID",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockPaymentUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&domain.Payment{}, nil).Times(1)
				m.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).Times(1)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: invalid card number": {
			input: input{
				ctx:        echo.New(),
				cardNumber: "123456781234567",
				expiryDate: "1122",
				cvv:        "123",
				userId:     "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockPaymentUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
		"failed: invalid expiry date": {
			input: input{
				ctx:        echo.New(),
				cardNumber: "1234567812345678",
				expiryDate: "112",
				cvv:        "123",
				userId:     "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockPaymentUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
		"failed: invalid cvv": {
			input: input{
				ctx:        echo.New(),
				cardNumber: "1234567812345678",
				expiryDate: "1122",
				cvv:        "12",
				userId:     "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockPaymentUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
		"failed: get payment": {
			input: input{
				ctx:        echo.New(),
				cardNumber: "1234567812345678",
				expiryDate: "1122",
				cvv:        "123",
				userId:     "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockPaymentUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to get Payment")).Times(1)
			},
		},
		"failed: add payment": {
			input: input{
				ctx:        echo.New(),
				cardNumber: "1234567812345678",
				expiryDate: "1122",
				cvv:        "123",
				userId:     "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockPaymentUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(xerrors.New("error")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			paymentUsecase := mock.NewMockPaymentUsecase(ctrl)
			tt.mockFunc(paymentUsecase)

			h := &paymentHandler{
				PaymentUsecase: paymentUsecase,
			}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(fmt.Sprintf(`{
				"cardNumber": "%s",
				"expiryDate": "%s",
				"cvv": "%s",
				"userId": "%s"
			}`, tt.input.cardNumber, tt.input.expiryDate, tt.input.cvv, tt.input.userId)))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)

			err := h.Add(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

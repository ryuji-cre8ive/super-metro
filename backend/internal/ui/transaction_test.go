package ui

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	mock "github.com/ryuji-cre8ive/super-metro/internal/usecase/mock"
	"go.uber.org/mock/gomock"
)

func TestTransactionHandler_Get(t *testing.T) {
	type input struct {
		ctx    *echo.Echo
		userId string
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockTransactionUsecase)
	}{
		"success": {
			input: input{
				ctx:    echo.New(),
				userId: "testID",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockTransactionUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return([]*domain.Transaction{}, nil).Times(1)
			},
		},
		"failed: get transaction": {
			input: input{
				ctx:    echo.New(),
				userId: "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockTransactionUsecase) {
				m.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to get Transaction")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			transactionUsecase := mock.NewMockTransactionUsecase(ctrl)
			tt.mockFunc(transactionUsecase)

			h := &transactionHandler{
				TransactionUsecase: transactionUsecase,
			}

			req := httptest.NewRequest(http.MethodGet, "/?userID="+tt.input.userId, nil)
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)

			err := h.Get(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

package usecase

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/stores"
	mock "github.com/ryuji-cre8ive/super-metro/internal/stores/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTransactionUsecase_Add(t *testing.T) {
	type input struct {
		ctx             echo.Context
		userId          string
		paymentId       string
		transactionType string
		amount          int
	}

	tests := map[string]struct {
		input    input
		want     error
		mockFunc func(m *mock.MockTransactionStore)
	}{
		"success": {
			input: input{
				ctx:             nil,
				userId:          "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				paymentId:       "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				transactionType: "TOPUP",
				amount:          100,
			},
			want: nil,
			mockFunc: func(m *mock.MockTransactionStore) {
				m.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockTransactionStore(ctrl)
			tt.mockFunc(store)

			u := &transactionUsecase{
				stores: &stores.Stores{
					Transaction: store,
				},
			}

			err := u.Add(tt.input.ctx, tt.input.userId, tt.input.paymentId, tt.input.transactionType, tt.input.amount)
			assert.Equal(t, tt.want, err)
		})
	}

	// ここにテストの実行部分を追加します
}

func TestTransactionUsecase_Get(t *testing.T) {
	type input struct {
		ctx    echo.Context
		userId string
	}

	tests := map[string]struct {
		input    input
		want     []*domain.Transaction
		wantErr  error
		mockFunc func(m *mock.MockTransactionStore)
	}{
		"success": {
			input: input{
				ctx:    nil,
				userId: "6d1c3e1b-d068-431f-b188-a436ac54ce52",
			},
			want: []*domain.Transaction{{
				UserID:          "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				PaymentID:       "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				TransactionType: "TOPUP",
				Amount:          100,
			}},
			wantErr: nil,
			mockFunc: func(m *mock.MockTransactionStore) {
				m.EXPECT().Get(gomock.Any()).Return([]*domain.Transaction{{
					UserID:          "6d1c3e1b-d068-431f-b188-a436ac54ce52",
					PaymentID:       "6d1c3e1b-d068-431f-b188-a436ac54ce52",
					TransactionType: "TOPUP",
					Amount:          100,
				}}, nil).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockTransactionStore(ctrl)
			tt.mockFunc(store)

			u := &transactionUsecase{
				stores: &stores.Stores{
					Transaction: store,
				},
			}

			transaction, err := u.Get(tt.input.ctx, tt.input.userId)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.want, transaction)
		})
	}
}

package usecase

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/stores"
	mock "github.com/ryuji-cre8ive/super-metro/internal/stores/mock"
	utilsMock "github.com/ryuji-cre8ive/super-metro/internal/utils/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/xerrors"
)

func TestUserUsecase_Create(t *testing.T) {
	type input struct {
		ctx      echo.Context
		email    string
		userName string
		password string
	}

	wantErr := xerrors.New("failed to encrypt password")

	tests := map[string]struct {
		input    input
		want     error
		mockFunc func(m *mock.MockUserStore, u *utilsMock.MockEncryptType)
	}{
		"success": {
			input: input{
				ctx:      nil,
				email:    "test@test.com",
				userName: "test",
				password: "test",
			},
			want: nil,
			mockFunc: func(m *mock.MockUserStore, u *utilsMock.MockEncryptType) {
				u.EXPECT().PasswordEncrypt("test").Return("encryptedTest", nil).Times(1)
				m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: empty-password": {
			input: input{
				ctx:      nil,
				email:    "test@test.com",
				userName: "test",
				password: "",
			},
			want: xerrors.Errorf("failed to encrypt password: %w", wantErr),
			mockFunc: func(m *mock.MockUserStore, u *utilsMock.MockEncryptType) {
				u.EXPECT().PasswordEncrypt(gomock.Any()).Return("", wantErr).Times(1)
				m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockUserStore(ctrl)
			encrypt := utilsMock.NewMockEncryptType(ctrl)
			tt.mockFunc(store, encrypt)

			u := &userUsecase{
				stores: &stores.Stores{
					User: store,
				},
				encrypt: encrypt,
			}

			err := u.Create(tt.input.ctx, tt.input.email, tt.input.userName, tt.input.password)
			if err != nil {
				assert.Equal(t, tt.want.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want, err)
			}
		})
	}
}

func TestUserUsecase_FindByEmail(t *testing.T) {
	type input struct {
		ctx   echo.Context
		email string
	}

	user := &domain.User{
		ID:       "6d1c3e1b-d068-431f-b188-a436ac54ce52",
		Email:    "test@test.com",
		Name:     "test",
		Password: "test",
		Valance:  100,
	}

	wantedError := xerrors.New("error")

	tests := map[string]struct {
		input    input
		want     *domain.User
		wantErr  error
		mockFunc func(m *mock.MockUserStore)
	}{
		"success": {
			input: input{
				ctx:   nil,
				email: "test@test.com",
			},
			want:    user,
			wantErr: nil,
			mockFunc: func(m *mock.MockUserStore) {
				m.EXPECT().FindByEmail(gomock.Any()).Return(user, nil).Times(1)
			},
		},
		"failed: standard-error": {
			input: input{
				ctx:   nil,
				email: "",
			},
			want:    nil,
			wantErr: xerrors.Errorf("failed to login: %w", wantedError),
			mockFunc: func(m *mock.MockUserStore) {
				m.EXPECT().FindByEmail(gomock.Any()).Return(nil, wantedError).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockUserStore(ctrl)
			tt.mockFunc(store)

			u := &userUsecase{
				stores: &stores.Stores{
					User: store,
				},
			}

			user, err := u.FindByEmail(tt.input.ctx, tt.input.email)
			if err != nil {
				assert.Equal(t, user, tt.want)
			} else {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestUserUsecase_TopUp(t *testing.T) {
	type input struct {
		ctx    echo.Context
		id     string
		amount int
	}

	wantedError := xerrors.New("error")

	tests := map[string]struct {
		input    input
		want     *domain.User
		wantErr  error
		mockFunc func(m *mock.MockUserStore)
	}{
		"success": {
			input: input{
				ctx:    nil,
				id:     "6d1c3e1b-d068-431f-b188-a436ac54ce52",
				amount: 100,
			},
			want:    &domain.User{},
			wantErr: nil,
			mockFunc: func(m *mock.MockUserStore) {
				m.EXPECT().TopUp(gomock.Any(), gomock.Any()).Return(&domain.User{}, nil).Times(1)
			},
		},
		"failed: standard-error": {
			input: input{
				ctx:    nil,
				id:     "",
				amount: 100,
			},
			want:    nil,
			wantErr: xerrors.Errorf("failed to top up: %w", wantedError),
			mockFunc: func(m *mock.MockUserStore) {
				m.EXPECT().TopUp(gomock.Any(), gomock.Any()).Return(nil, wantedError).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockUserStore(ctrl)
			tt.mockFunc(store)

			u := &userUsecase{
				stores: &stores.Stores{
					User: store,
				},
			}

			user, err := u.TopUp(tt.input.ctx, tt.input.id, tt.input.amount)
			if err != nil {
				assert.Equal(t, user, tt.want)
			} else {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestUserUsecase_GetSession(t *testing.T) {
	type input struct {
		ctx echo.Context
		id  string
	}

	wantedError := xerrors.New("error")

	tests := map[string]struct {
		input    input
		want     string
		wantErr  error
		mockFunc func(m *mock.MockUserStore)
	}{
		"success": {
			input: input{
				ctx: nil,
				id:  "6d1c3e1b-d068-431f-b188-a436ac54ce52",
			},
			want:    "",
			wantErr: nil,
			mockFunc: func(m *mock.MockUserStore) {
				m.EXPECT().GetSession(gomock.Any()).Return("", nil).Times(1)
			},
		},
		"failed: standard-error": {
			input: input{
				ctx: nil,
				id:  "",
			},
			want:    "",
			wantErr: xerrors.Errorf("failed to get session: %w", wantedError),
			mockFunc: func(m *mock.MockUserStore) {
				m.EXPECT().GetSession(gomock.Any()).Return("", wantedError).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock.NewMockUserStore(ctrl)
			tt.mockFunc(store)

			u := &userUsecase{
				stores: &stores.Stores{
					User: store,
				},
			}

			session, err := u.GetSession(tt.input.ctx, tt.input.id)
			if err != nil {
				assert.Equal(t, session, tt.want)
			} else {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

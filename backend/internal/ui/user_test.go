package ui

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	mock "github.com/ryuji-cre8ive/super-metro/internal/usecase/mock"
	"golang.org/x/xerrors"

	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestUserHandler_Create(t *testing.T) {
	type input struct {
		ctx      *echo.Echo
		email    string
		userName string
		password string
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockUserUsecase)
	}{
		"success": {
			input: input{
				ctx:      echo.New(),
				email:    "test@test.com",
				userName: "test",
				password: "test",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: create user": {
			input: input{
				ctx:      echo.New(),
				email:    "test@test.com",
				userName: "test",
				password: "test",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("failed to post User")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			tt.mockFunc(userUsecase)

			h := &userHandler{
				UserUsecase: userUsecase,
			}

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)

			err := h.Create(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserHandler_Login(t *testing.T) {
	type input struct {
		ctx      *echo.Echo
		email    string
		password string
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockUserUsecase)
	}{
		"success": {
			input: input{
				ctx:      echo.New(),
				email:    "test@test.com",
				password: "test",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(&domain.User{}, nil).Times(1)
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: find by email": {
			input: input{
				ctx:      echo.New(),
				email:    "test@test.com",
				password: "test",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to login")).Times(1)
			},
		},
		"failed: set session": {
			input: input{
				ctx:      echo.New(),
				email:    "test@test.com",
				password: "test",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(&domain.User{}, nil).Times(1)
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(xerrors.New("set session error")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			tt.mockFunc(userUsecase)

			h := &userHandler{
				UserUsecase: userUsecase,
			}

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)

			err := h.Login(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserHandler_Logout(t *testing.T) {
	type input struct {
		ctx    *echo.Echo
		userID string
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockUserUsecase)
	}{
		"success": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: set session": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("failed to set session")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			tt.mockFunc(userUsecase)

			h := &userHandler{
				UserUsecase: userUsecase,
			}

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)

			err := h.Logout(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserHandler_TopUp(t *testing.T) {
	type input struct {
		ctx    *echo.Echo
		userID string
		amount int
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockUserUsecase, p *mock.MockPaymentUsecase, t *mock.MockTransactionUsecase)
	}{
		"success": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
				amount: 100,
			},
			wantErr: false,
			mockFunc: func(m *mock.MockUserUsecase, p *mock.MockPaymentUsecase, t *mock.MockTransactionUsecase) {
				m.EXPECT().TopUp(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.User{}, nil).Times(1)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&domain.Payment{}, nil).Times(1)
				t.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: top up": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
				amount: 100,
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase, p *mock.MockPaymentUsecase, t *mock.MockTransactionUsecase) {
				m.EXPECT().TopUp(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("failed to top up")).Times(1)
			},
		},
		"failed: payment error": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
				amount: 100,
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase, p *mock.MockPaymentUsecase, t *mock.MockTransactionUsecase) {
				m.EXPECT().TopUp(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.User{}, nil).Times(1)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, xerrors.New("get payment error")).Times(1)
				t.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
		"failed: transaction error": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
				amount: 100,
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase, p *mock.MockPaymentUsecase, t *mock.MockTransactionUsecase) {
				m.EXPECT().TopUp(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.User{}, nil).Times(1)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&domain.Payment{}, nil).Times(1)
				t.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(xerrors.New("transaction error")).Times(1)
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
		"failed: set session error": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
				amount: 100,
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase, p *mock.MockPaymentUsecase, t *mock.MockTransactionUsecase) {
				m.EXPECT().TopUp(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.User{}, nil).Times(1)
				p.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&domain.Payment{}, nil).Times(1)
				t.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
				m.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(xerrors.New("set session error")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			paymentUsecase := mock.NewMockPaymentUsecase(ctrl)
			transactionUsecase := mock.NewMockTransactionUsecase(ctrl)
			tt.mockFunc(userUsecase, paymentUsecase, transactionUsecase)

			h := &userHandler{
				UserUsecase:        userUsecase,
				PaymentUsecase:     paymentUsecase,
				TransactionUsecase: transactionUsecase,
			}

			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)

			err := h.TopUp(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserHandler_Cookie(t *testing.T) {
	type input struct {
		ctx    *echo.Echo
		userID string
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockUserUsecase)
	}{
		"success": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().GetSession(gomock.Any(), gomock.Any()).Return("sessionToken", nil).Times(1)
			},
		},
		"failed: get session": {
			input: input{
				ctx:    echo.New(),
				userID: "testID",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().GetSession(gomock.Any(), gomock.Any()).Return("", errors.New("failed to get session")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			tt.mockFunc(userUsecase)

			h := &userHandler{
				UserUsecase: userUsecase,
			}

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)
			c.SetParamNames("userID")
			c.SetParamValues(tt.input.userID)

			err := h.Cookie(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserHandler_CheckCookieExpiration(t *testing.T) {
	tests := map[string]struct {
		cookie   *http.Cookie
		wantErr  bool
		mockFunc func(m *mock.MockUserUsecase)
	}{
		"success": {
			cookie: &http.Cookie{
				Name:  "session_token",
				Value: "sessionToken",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().IsCookieExist(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: cookie does not exist": {
			cookie: &http.Cookie{
				Name:  "session_token",
				Value: "sessionToken",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().IsCookieExist(gomock.Any(), gomock.Any()).Return(errors.New("Need to Authorize")).Times(1)
			},
		},
		"failed: cookie name does not session_token": {
			cookie: &http.Cookie{
				Name:  "token",
				Value: "sessionToken",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().IsCookieExist(gomock.Any(), gomock.Any()).Return(nil).Times(0)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			tt.mockFunc(userUsecase)

			h := &userHandler{
				UserUsecase: userUsecase,
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.AddCookie(tt.cookie)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			middlewareFunc := h.CheckCookieExpiration()
			err := middlewareFunc(func(c echo.Context) error {
				return nil
			})(c)

			if (err != nil) != tt.wantErr {
				t.Errorf("CheckCookieExpiration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserHandler_IsCookieExist(t *testing.T) {
	type input struct {
		ctx         *echo.Echo
		cookieValue string
	}

	tests := map[string]struct {
		input    input
		wantErr  bool
		mockFunc func(m *mock.MockUserUsecase)
	}{
		"success": {
			input: input{
				ctx:         echo.New(),
				cookieValue: "sessionToken",
			},
			wantErr: false,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().IsCookieExist(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
		},
		"failed: cookie does not exist": {
			input: input{
				ctx:         echo.New(),
				cookieValue: "sessionToken",
			},
			wantErr: true,
			mockFunc: func(m *mock.MockUserUsecase) {
				m.EXPECT().IsCookieExist(gomock.Any(), gomock.Any()).Return(errors.New("Need to Authorize")).Times(1)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userUsecase := mock.NewMockUserUsecase(ctrl)
			tt.mockFunc(userUsecase)

			h := &userHandler{
				UserUsecase: userUsecase,
			}

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := tt.input.ctx.NewContext(req, rec)

			err := h.IsCookieExist(c, tt.input.cookieValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsCookieExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

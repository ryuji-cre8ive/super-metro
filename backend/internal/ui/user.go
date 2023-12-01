package ui

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/usecase"
	"github.com/ryuji-cre8ive/super-metro/internal/utils"
	"golang.org/x/xerrors"
)

type (
	UserHandler interface {
		Create(c echo.Context) error
		Login(c echo.Context) error
		Logout(c echo.Context) error
		TopUp(c echo.Context) error
	}

	userHandler struct {
		usecase.UserUsecase
		usecase.PaymentUsecase
		usecase.TransactionUsecase
	}
)

func (h *userHandler) Create(c echo.Context) error {
	// var user domain.User
	param := new(domain.User)
	if err := c.Bind(param); err != nil {
		return xerrors.Errorf("failed to bind User: %w", err)
	}
	name := param.Name
	password := param.Password
	email := param.Email
	createErr := h.UserUsecase.Create(c, email, name, password)
	if createErr != nil {
		return xerrors.Errorf("failed to post User: %w", createErr)
	}

	return c.String(200, "success")
}

func (h *userHandler) Login(c echo.Context) error {
	param := new(domain.User)
	if err := c.Bind(param); err != nil {
		return xerrors.Errorf("failed to bind User: %w", err)
	}
	password := param.Password
	email := param.Email
	user, loginErr := h.UserUsecase.FindByEmail(c, email)
	if loginErr != nil {
		return xerrors.Errorf("failed to login: %w", loginErr)
	}
	compareErr := utils.CheckHashPassword(user.Password, password)
	if compareErr != nil {
		xerrors.Errorf("failed to compare password: %w", compareErr)
	}

	// JWT認証の作成
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["userName"] = user.Name
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["valance"] = user.Valance

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return xerrors.Errorf("failed to create JWT: %w", err)
	}
	c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+t)

	user.SessionToken = t

	return c.JSON(200, user)
}

// This function handles user logout, invalidates the JWT by setting its expiration to the past, and sets this JWT in the response header
func (h *userHandler) Logout(c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = ""
	claims["exp"] = time.Now().Add(-time.Hour).Unix()

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return xerrors.Errorf("failed to create JWT: %w", err)
	}
	c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+t)

	return c.String(200, "success")
}

func (h *userHandler) TopUp(c echo.Context) error {
	const TOP_UP_TRANSACTION_TYPE = "TOPUP"
	param := new(domain.User)
	if err := c.Bind(param); err != nil {
		return xerrors.Errorf("failed to bind User: %w", err)
	}
	id := param.ID
	amount := param.Valance

	user, topUpErr := h.UserUsecase.TopUp(c, id, amount)
	if topUpErr != nil {
		return xerrors.Errorf("failed to top up: %w", topUpErr)
	}
	userPaymentInfo, paymentErr := h.PaymentUsecase.Get(c, id)
	if paymentErr != nil {
		return xerrors.Errorf("failed to get Payment: %w", paymentErr)
	}
	transactionErr := h.TransactionUsecase.Add(c, id, userPaymentInfo.ID, TOP_UP_TRANSACTION_TYPE, amount)
	if transactionErr != nil {
		return xerrors.Errorf("failed to post Transaction: %w", transactionErr)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["userName"] = user.Name
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["valance"] = user.Valance

	t, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return xerrors.Errorf("failed to create JWT: %w", err)
	}
	c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+t)
	return c.JSON(200, map[string]interface{}{"sessionToken": t})
}

package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"github.com/ryuji-cre8ive/super-metro/internal/stores"
	"github.com/ryuji-cre8ive/super-metro/internal/utils"
	"golang.org/x/xerrors"
)

type (
	UserUsecase interface {
		Create(ctx echo.Context, email string, userName string, password string) error
		FindByEmail(ctx echo.Context, email string) (*domain.User, error)
		TopUp(ctx echo.Context, id string, amount int) (*domain.User, error)
		GetSession(ctx echo.Context, id string) (string, error)
		SetSession(ctx echo.Context, id string, session string) error
		IsCookieExist(ctx echo.Context, cookieValue string) error
	}

	userUsecase struct {
		stores  *stores.Stores
		encrypt utils.EncryptType
	}
)

func (u *userUsecase) Create(ctx echo.Context, email string, userName string, password string) error {
	encryptedPassword, err := u.encrypt.PasswordEncrypt(password)
	if err != nil {
		return xerrors.Errorf("failed to encrypt password: %w", err)
	}
	return u.stores.User.Create(email, userName, encryptedPassword)
}

func (u *userUsecase) FindByEmail(ctx echo.Context, email string) (*domain.User, error) {
	user, loginErr := u.stores.User.FindByEmail(email)
	if loginErr != nil {
		return nil, xerrors.Errorf("failed to login: %w", loginErr)
	}
	return user, nil
}

func (u *userUsecase) TopUp(ctx echo.Context, id string, amount int) (*domain.User, error) {
	user, err := u.stores.User.TopUp(id, amount)
	if err != nil {
		return nil, xerrors.Errorf("failed to top up: %w", err)
	}
	return user, nil
}

func (u *userUsecase) GetSession(ctx echo.Context, id string) (string, error) {
	return u.stores.User.GetSession(id)
}

func (u *userUsecase) SetSession(ctx echo.Context, id string, session string) error {
	return u.stores.User.SetSession(id, session)
}

func (u *userUsecase) IsCookieExist(ctx echo.Context, cookieValue string) error {
	return u.stores.User.IsCookieExist(cookieValue)
}

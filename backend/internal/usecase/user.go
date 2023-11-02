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
	}

	userUsecase struct {
		stores *stores.Stores
	}
)

func (u *userUsecase) Create(ctx echo.Context, email string, userName string, password string) error {
	encryptedPassword, err := utils.PasswordEncrypt(password)
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

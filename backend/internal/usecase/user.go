package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-suica/internal/stores"
	"github.com/ryuji-cre8ive/super-suica/internal/utils"
	"golang.org/x/xerrors"
)

type (
	UserUsecase interface {
		Create(ctx echo.Context, userName string, password string) error
	}

	userUsecase struct {
		stores *stores.Stores
	}
)

func (u *userUsecase) Create(ctx echo.Context, userName string, password string) error {
	encryptedPassword, err := utils.PasswordEncrypt(password)
	if err != nil {
		return xerrors.Errorf("failed to encrypt password: %w", err)
	}
	return u.stores.User.Create(userName, encryptedPassword)
}

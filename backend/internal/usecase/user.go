package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/super-suica/internal/stores"
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
	return u.stores.User.Create(userName, password)
}

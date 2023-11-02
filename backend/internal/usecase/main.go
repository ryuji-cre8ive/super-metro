package usecase

import "github.com/ryuji-cre8ive/super-metro/internal/stores"

type Usecase struct {
	User UserUsecase
}

func New(s *stores.Stores) *Usecase {
	return &Usecase{
		User: &userUsecase{stores: s},
	}
}

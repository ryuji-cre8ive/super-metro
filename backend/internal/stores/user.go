package stores

import (
	"github.com/ryuji-cre8ive/super-suica/internal/domain"
	"gorm.io/gorm"
)

type (
	UserStore interface {
		Create(userName string, password string) error
	}

	userStore struct {
		*gorm.DB
	}
)

func (s *userStore) Create(userName string, password string) error {
	return s.DB.Create(&domain.User{
		Name:     userName,
		Password: password,
	}).Error
}

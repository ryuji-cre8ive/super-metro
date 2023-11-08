package stores

import (
	"time"

	"github.com/google/uuid"
	"github.com/ryuji-cre8ive/super-metro/internal/domain"
	"golang.org/x/xerrors"

	"gorm.io/gorm"
)

type (
	UserStore interface {
		Create(email string, userName string, password string) error
		FindByEmail(email string) (*domain.User, error)
		TopUp(id string, amount int) (*domain.User, error)
	}

	userStore struct {
		*gorm.DB
	}
)

func (s *userStore) Create(email string, userName string, password string) error {
	uuid := uuid.Must(uuid.NewRandom())
	return s.DB.Create(&domain.User{
		ID:        uuid.String(),
		Email:     email,
		Name:      userName,
		Password:  password,
		Valance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}

func (s *userStore) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := s.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, xerrors.Errorf("cannot find user by email: %w", result.Error)
	}

	return &user, nil
}

func (s *userStore) TopUp(id string, amount int) (*domain.User, error) {
	var user domain.User
	result := s.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, xerrors.Errorf("cannot find user by id: %w", result.Error)
	}
	user.Valance += amount
	return &user, s.DB.Save(&user).Error
}

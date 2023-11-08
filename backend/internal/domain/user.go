package domain

import (
	"time"
)

type User struct {
	ID           string     `json:"id"`
	Name         string     `json:"userName"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	SessionToken string     `json:"sessionToken"`
	Valance      int        `json:"valance"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
}

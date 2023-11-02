package domain

type User struct {
	ID           string `json:"id"`
	Name         string `json:"userName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	SessionToken string `json:"sessionToken"`
}

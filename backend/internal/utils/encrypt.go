package utils

import "golang.org/x/crypto/bcrypt"

// 暗号化 (hash)
func PasswordEncrypt(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword), err
}

// 暗号化パスワードと比較
func CheckHashPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

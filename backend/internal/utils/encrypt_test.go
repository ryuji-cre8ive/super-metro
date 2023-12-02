package utils

import (
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	password := "testPassword"
	hashedPassword, err := PasswordEncrypt(password)
	if err != nil {
		t.Errorf("PasswordEncrypt failed with error: %v", err)
	}

	err = CheckHashPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("CheckHashPassword failed with error: %v", err)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("secretsecretsecretsecret")
	plainText := []byte("this is some text to encrypt")

	encrypted, err := Encrypt(plainText, key)
	if err != nil {
		t.Errorf("Encrypt failed with error: %v", err)
	}

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Errorf("Decrypt failed with error: %v", err)
	}

	if string(decrypted) != string(plainText) {
		t.Errorf("Decrypt text doesn't match original text")
	}
}
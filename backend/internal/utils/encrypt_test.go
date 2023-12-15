package utils

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncrypt(t *testing.T) {
	type input struct {
		text []byte
		key  []byte
	}

	tests := map[string]struct {
		input input
		want  error
	}{
		"success": {
			input: input{
				text: []byte("testPassword"),
				key:  []byte("secretsecretsecretsecret"),
			},
			want: nil,
		},
		"failed: wrong-key": {
			input: input{
				text: []byte("testPassword"),
				key:  []byte("secret"),
			},
			want: aes.KeySizeError(6),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := Encrypt(tt.input.text, tt.input.key)
			assert.Equal(t, tt.want, err)
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	type input struct {
		plainText []byte
		key       []byte
	}

	tests := map[string]struct {
		input input
		want  string
		err   error
	}{
		"success": {
			input: input{
				plainText: []byte("this is some text to encrypt"),
				key:       []byte("secretsecretsecretsecret"),
			},
			want: "this is some text to encrypt",
			err:  nil,
		},
		"failed: wrong-key": {
			input: input{
				plainText: []byte("this is some text to encrypt"),
				key:       []byte("secret"),
			},
			want: "",
			err:  aes.KeySizeError(6),
		},
		"failed: empty-text": {
			input: input{
				plainText: []byte(""),
				key:       []byte("secretsecretsecretsecret"),
			},
			want: "",
			err:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			encrypted, err := Encrypt(tt.input.plainText, tt.input.key)
			if err != tt.err {
				t.Errorf("Encrypt failed with error: %v", err)
			}

			decrypted, err := Decrypt(encrypted, tt.input.key)
			if err != tt.err {
				t.Errorf("Decrypt failed with error: %v", err)
			}

			if string(decrypted) != tt.want {
				t.Errorf("Decrypt text doesn't match original text")
			}
		})
	}
}

func TestPasswordEncrypt(t *testing.T) {
	type input struct {
		password string
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				password: "testPassword",
			},
			wantErr: false,
		},
		"empty password": {
			input: input{
				password: "",
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			hashedPassword, err := PasswordEncrypt(tt.input.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswordEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(tt.input.password))
				if err != nil {
					t.Errorf("PasswordEncrypt() failed password verification, error: %v", err)
				}
			}
		})
	}
}

func TestCheckHashPassword(t *testing.T) {
	type input struct {
		password string
	}

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"success": {
			input: input{
				password: "testPassword",
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			hashedPassword, _ := PasswordEncrypt(tt.input.password)
			err := CheckHashPassword(hashedPassword, tt.input.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckHashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

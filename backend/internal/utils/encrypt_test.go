package utils

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordEncrypt(t *testing.T) {
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

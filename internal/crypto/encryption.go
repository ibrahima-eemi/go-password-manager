package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Encrypt chiffre un texte avec AES-GCM
func Encrypt(data, key string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("clé de chiffrement invalide : doit être de 32 caractères")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt déchiffre un texte chiffré avec AES-GCM
func Decrypt(encryptedData, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	if len(key) != 32 {
		return "", errors.New("clé de chiffrement invalide : doit être de 32 caractères")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(data) < aesGCM.NonceSize() {
		return "", errors.New("données invalides")
	}

	nonce := data[:aesGCM.NonceSize()]
	cipherText := data[aesGCM.NonceSize():]

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword génère un hash sécurisé
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHash vérifie un mot de passe contre son hash
func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword returns the bcrypt hash of a plain-text password.
func HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword compares a plain-text password against a bcrypt hash.
func CheckPassword(hash, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}

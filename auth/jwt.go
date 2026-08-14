package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims is the JWT payload used across the application.
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	RoleName string `json:"role_name"`
	jwt.RegisteredClaims
}

// GenerateToken creates a signed JWT for the given user identity.
func GenerateToken(secret string, expireHours int, userID int, username, roleName string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RoleName: roleName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken validates a JWT string and returns its claims.
func ParseToken(secret, tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth: unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("auth: invalid token")
	}
	return claims, nil
}

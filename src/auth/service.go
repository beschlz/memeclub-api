package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func AuthorizeUser(creds *Credentials) (string, error) {
	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		return "", fmt.Errorf("unauthorized")
	}

	expirationTime := jwt.NumericDate{Time: time.Now().Add(time.Minute * 30)}

	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)

	return tokenString, err
}

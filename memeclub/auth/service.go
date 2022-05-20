package auth

import (
	"fmt"
	"github.com/beschlz/memeclub-api/memeclub/users"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func AuthorizeUser(creds *Credentials) (string, error) {
	user, err := users.GetUserBayName(creds.Username)

	if err != nil {
		return "", err
	}

	authOK := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))

	if authOK != nil {
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

func ValidateToken(token string) error {
	claims := &Claims{}

	parsedToken, err := jwt.ParseWithClaims(
		token,
		claims,
		func(jwtToken *jwt.Token) (interface{}, error) {
			return key, nil
		})

	if err != nil || !parsedToken.Valid {
		return err
	}

	return nil
}

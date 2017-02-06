package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims that defines the JWT data
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken Generate jwt new token for the user
func (c Claims) CreateToken(username string) (string, int64, error) {
	// Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	// We'll manually assign the claims but in production you'd insert values from a database
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, err := token.SignedString([]byte("secret")) // Take Secret From Config

	return signedToken, expireToken, err
}

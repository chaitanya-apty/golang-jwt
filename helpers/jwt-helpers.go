package helpers

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type AuthClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

//GenerateAuthToken - Generates Auth Session
func GenerateAuthToken(secret []byte) (string, error) {
	//Creating Claims
	claims := AuthClaims{
		"Chaitanya Kumar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "Chai-Server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("Cannot Generate Token %s", err.Error())
	}
	return tokenString, nil
}

//ValidateSession - Validates Auth Session
func ValidateSession(sessionToken string, secret []byte) (string, error) {
	token, err := jwt.ParseWithClaims(sessionToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err == nil && token.Valid {
		claims, _ := token.Claims.(*AuthClaims)
		return "Session Verified " + claims.User, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", fmt.Errorf("Malformed Token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "", fmt.Errorf("Token is Expired")
		} else {
			return "", fmt.Errorf("Error: %s", err.Error())
		}
	} else {
		return "Session Error", err
	}
}

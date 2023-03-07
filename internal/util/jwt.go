package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

const jwtKey = "secret"

func GenerateJWT(email, username string) (string, error) {
	expTime := time.Now().Add(30 * time.Minute)
	claims := JWTClaim{Username: username, Email: email,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expTime.Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := parsedToken.Claims.(*JWTClaim)
	if !ok {
		return err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return err
	}
	return nil
}

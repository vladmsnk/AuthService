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

func GenerateJWT(email, username, signingKey string, timeToLive int) (string, error) {
	expTime := time.Now().Add(time.Duration(timeToLive) * time.Second)
	claims := JWTClaim{Username: username, Email: email, StandardClaims: jwt.StandardClaims{ExpiresAt: expTime.Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token, signingKey string) error {

	parsedTkn, err := jwt.ParseWithClaims(token, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningAlgorithm
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return ErrInvalidAccessToken
	}

	claims, ok := parsedTkn.Claims.(*JWTClaim)
	if !ok || !parsedTkn.Valid {
		return ErrInvalidAccessToken
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return ErrExpiredToken
	}
	return nil
}

package middleware

import (
	"errors"
	"time"

	"github.com/NetSinx/go-restful-API/model/domain"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var myKey = []byte("yasinganteng")

func GenerateToken() string {
	var user domain.Login

	claims := &JWTClaim{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, _ := token.SignedString(myKey)

	return tokenString
}

func ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		err = errors.New("token is expired")
		return err
	}

	return nil
}
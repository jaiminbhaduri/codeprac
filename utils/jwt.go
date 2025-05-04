package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = os.Getenv("SECRET_KEY")
var jwtKey = []byte(secretkey)

type Claims struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string, id int) (string, error) {
	jwt_expire_hours_str := os.Getenv("JWT_EXPIRE_HOURS")
	jwt_expire_hours, _ := strconv.Atoi(jwt_expire_hours_str)

	expirationTime := time.Now().Add(time.Duration(jwt_expire_hours) * time.Hour)
	claims := &Claims{
		Username: username,
		Id:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

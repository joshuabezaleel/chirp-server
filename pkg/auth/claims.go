package auth

import (
	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("secretkey")

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

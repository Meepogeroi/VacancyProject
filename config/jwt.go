package config

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	Mail string `json:"mail"`
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

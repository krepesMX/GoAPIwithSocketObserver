package models

import "github.com/golang-jwt/jwt"

type AppClaims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

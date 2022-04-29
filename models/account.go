package models

import "github.com/golang-jwt/jwt"

type Account struct {
	Login       string `json:"login" db:"login"`
	Password    string `json:"-" db:"password"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	Permissions string
}

type JwtClaims struct {
	jwt.StandardClaims
	Permissions string
}

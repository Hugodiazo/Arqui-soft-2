package utils

import (
	"github.com/golang-jwt/jwt"
)

// Claims define las claims para el token JWT
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

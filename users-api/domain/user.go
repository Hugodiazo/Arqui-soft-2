package domain

import "github.com/golang-jwt/jwt"

// Estructura del usuario
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

// Estructura de credenciales para inicio de sesión
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Estructura de las claims para el token JWT
type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

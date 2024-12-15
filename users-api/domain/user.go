package domain

import "github.com/golang-jwt/jwt"

// Estructura del usuario
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

// Estructura de credenciales para inicio de sesión
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims representa los claims del token JWT
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

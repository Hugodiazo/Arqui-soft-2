// cursos-api/utils/jwt.go

package utils

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("tu_clave_secreta")

func ExtractUserIDFromToken(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("No token provided")
	}

	// Elimina el prefijo "Bearer "
	tokenString = tokenString[len("Bearer "):]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["user_id"].(string), nil
	} else {
		return "", errors.New("Invalid token")
	}
}

package services

import (
	"errors"
	"time"

	"users-app/users-api/dao"
	"users-app/users-api/domain"
	"users-app/users-api/utils"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your_secret_key")

// Registrar un nuevo usuario
func RegisterUser(user domain.User) error {
	existingUser, _ := dao.GetUserByEmail(user.Email)
	if existingUser.ID != 0 {
		return errors.New("usuario ya existe")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return dao.CreateUser(user)
}

// Iniciar sesión y generar JWT con rol
func LoginUser(credentials domain.Credentials) (string, error) {
	// Buscar el usuario por email
	user, err := dao.GetUserByEmail(credentials.Email)
	if err != nil || !utils.CheckPasswordHash(credentials.Password, user.Password) {
		return "", errors.New("credenciales inválidas")
	}

	// Crear el token JWT con Claims, incluyendo el rol
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &domain.Claims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Obtener todos los usuarios
func GetAllUsers() ([]domain.User, error) {
	return dao.GetAllUsers()
}

// Obtener un usuario por ID
func GetUserByID(userID int) (domain.User, error) {
	return dao.GetUserByID(userID)
}

// Actualizar un usuario
func UpdateUser(userID int, user domain.User) error {
	// Hashear la contraseña si se está actualizando
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	return dao.UpdateUser(userID, user)
}

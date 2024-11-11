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
	// Verifica si el usuario ya existe
	existingUser, _ := dao.GetUserByEmail(user.Email)
	if existingUser.ID != 0 {
		return errors.New("usuario ya existe")
	}

	// Hashear la contraseña antes de guardarla
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Insertar el usuario en la base de datos
	return dao.CreateUser(user)
}

// Iniciar sesión y generar JWT
func LoginUser(credentials domain.Credentials) (string, error) {
	// Buscar el usuario por email
	user, err := dao.GetUserByEmail(credentials.Email)
	if err != nil || !utils.CheckPasswordHash(credentials.Password, user.Password) {
		return "", errors.New("credenciales inválidas")
	}

	// Crear el token JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &domain.JWTClaims{
		UserID: user.ID,
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

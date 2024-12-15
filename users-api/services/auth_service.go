package services

import (
	"errors"
	"log"
	"time"

	"users-app/users-api/dao"
	"users-app/users-api/domain"
	"users-app/users-api/utils"

	"os"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

// Registrar un nuevo usuario
func RegisterUser(user domain.User) error {
	existingUser, _ := dao.GetUserByEmail(user.Email)
	if existingUser.ID != 0 {
		return errors.New("usuario ya existe")
	}

	// No es necesario hashear aquí, ya se hizo en el controlador
	return dao.CreateUser(user)
}

// Iniciar sesión y generar JWT con rol
func LoginUser(credentials domain.Credentials) (string, error) {
	log.Println("Buscando usuario por email:", credentials.Email)

	user, err := dao.GetUserByEmail(credentials.Email)
	if err != nil {
		log.Println("Error al obtener usuario:", err)
		return "", errors.New("usuario no encontrado")
	}

	log.Println("Usuario encontrado:", user.Email)
	log.Println("Comparando contraseñas")

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		log.Println("La contraseña no coincide")
		return "", errors.New("credenciales inválidas")
	}

	log.Println("Contraseña válida. Generando token.")

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &domain.Claims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(utils.SecretKey)
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

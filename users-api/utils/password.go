package utils

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// Genera un hash de la contraseña
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// Verifica una contraseña con su hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Genera y muestra una contraseña hasheada (para uso manual o depuración)
func GenerateHashedPassword(password string) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		fmt.Println("Error al hashear la contraseña:", err)
		return
	}
	fmt.Println("Contraseña hasheada:", hashedPassword)
}

// Cargar la clave secreta desde una variable de entorno
var SecretKey []byte

func init() {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		log.Fatal("SECRET_KEY no está configurado")
	}
	SecretKey = []byte(key)
}

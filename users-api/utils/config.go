package utils

import (
	"log"
	"os"
)

var SecretKey []byte

func init() {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		log.Fatal("SECRET_KEY no está configurado")
	}
	SecretKey = []byte(key)
}

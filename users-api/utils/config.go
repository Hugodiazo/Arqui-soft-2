package utils

import (
	"log"
	"os"
)

func init() {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		log.Fatal("SECRET_KEY no está configurado")
	}
	SecretKey = []byte(key)
}

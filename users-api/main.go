package main

import (
	"log"
	"net/http"
	"os"
	"users-app/users-api/db"
	"users-app/users-api/router"
	"users-app/users-api/utils"
)

// Middleware para habilitar CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Verificar si SECRET_KEY está configurado
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY no está configurado")
	}
	utils.SecretKey = []byte(secretKey)

	// Conectar a la base de datos
	db.ConnectDB()

	// Conectar a Memcached
	db.ConnectCache()

	// Configurar las rutas
	r := router.InitRoutes()

	// Aplicar el middleware de CORS
	handler := enableCors(r)

	log.Println("API de usuarios iniciada en http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

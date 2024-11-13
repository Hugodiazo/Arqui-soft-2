// cursos-api/main.go
package main

import (
	"log"
	"net/http"

	"cursos-app/cursos-api/clients"
	"cursos-app/cursos-api/db"
	"cursos-app/cursos-api/router"
)

// Middleware para habilitar CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
	// Conexión a MongoDB
	db.ConnectMongoDB("mongodb://localhost:27017", "arqsoft2")
	// Conexión a RabbitMQ
	clients.ConnectRabbitMQ()

	// Inicializar las rutas del router
	r := router.InitRoutes()

	// Aplicar middleware CORS a las rutas
	handler := enableCors(r)

	log.Println("API de cursos iniciada en http://localhost:8081")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatal(err)
	}
}

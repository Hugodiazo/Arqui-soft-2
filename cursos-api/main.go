// cursos-api/main.go
package main

import (
	"log"
	"net/http"
	"os"

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
	// Obtener la URI y el nombre de la base de datos desde variables de entorno
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	// Conexión a MongoDB
	db.ConnectMongoDB(mongoURI, dbName)

	// Conexión a RabbitMQ
	clients.ConnectRabbitMQ()

	// Inicializar las rutas del router
	r := router.InitRoutes()

	// Aplicar middleware CORS a las rutas
	handler := enableCors(r)

	// Iniciar el servidor en el puerto 8080
	log.Println("API de cursos iniciada en http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

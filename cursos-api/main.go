// cursos-api/main.go
package main

import (
	"log"
	"net/http"

	"cursos-app/cursos-api/clients"
	"cursos-app/cursos-api/db"
	"cursos-app/cursos-api/router"
)

func main() {
	// Conexión a MongoDB
	db.ConnectMongoDB("mongodb://localhost:27017", "arqsoft2")
	clients.ConnectRabbitMQ()

	// Inicializar las rutas
	r := router.InitRoutes()

	log.Println("API de cursos iniciada en http://localhost:8081")
	http.ListenAndServe(":8081", r)
}

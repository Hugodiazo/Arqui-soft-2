package main

import (
	"log"
	"net/http"

	"search-app/search-api/clients"
	"search-app/search-api/router"
)

func main() {
	clients.ConnectRabbitMQ() // Conexión a RabbitMQ
	r := router.InitRoutes()  // Inicializar las rutas

	log.Println("API de búsqueda iniciada en http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}

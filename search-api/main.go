// search-api/main.go
package main

import (
	"log"
	"net/http"

	"search-api/router"
	"search-api/services"
	"search-api/utils"
)

func main() {
	utils.ConnectRabbitMQ()
	defer utils.CloseRabbitMQ()

	go services.ProcessCourseUpdates() // Procesar actualizaciones de cursos

	r := router.InitRoutes()

	log.Println("API de búsqueda iniciada en http://localhost:8083")
	http.ListenAndServe(":8083", r)
}

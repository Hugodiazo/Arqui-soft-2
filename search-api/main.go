// search-api/main.go
package main

import (
	"log"
	"net/http"
	"search-app/search-api/router"
	"search-app/search-api/services"
	"search-app/search-api/utils"
)

func main() {
	// Conectar a RabbitMQ
	utils.ConnectRabbitMQ()
	defer utils.CloseRabbitMQ()

	// Suscribirse a la cola de actualizaciones de cursos
	err := utils.SubscribeToQueue("course_updates", services.HandleRabbitMQMessage)
	if err != nil {
		log.Fatalf("Error al suscribirse a la cola de RabbitMQ: %v", err)
	}

	// Iniciar el servidor de la API de búsqueda
	r := router.InitRoutes()
	log.Println("API de búsqueda iniciada en http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}

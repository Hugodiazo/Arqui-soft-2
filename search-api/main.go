package main

import (
	"log"
	"net/http"
	"search-app/search-api/router"
	"search-app/search-api/services"
	"search-app/search-api/utils"

	"github.com/gorilla/handlers"
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

	// Iniciar el servidor de la API de búsqueda con CORS habilitado
	r := router.InitRoutes()

	// Configurar CORS permitiendo solicitudes desde localhost:3000
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),         // Permitir solicitudes desde el frontend
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),  // Métodos HTTP permitidos
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Headers permitidos
	)

	log.Println("API de búsqueda iniciada en http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", corsOptions(r)))
}

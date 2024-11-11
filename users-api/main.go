// users-api/main.go
package main

import (
	"log"
	"net/http"
	"users-app/users-api/db"
	"users-app/users-api/router"
)

func main() {
	// Conectar a la base de datos
	db.ConnectDB()

	// Configurar las rutas
	r := router.InitRoutes()

	log.Println("API de usuarios iniciada en http://localhost:8082")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatal(err)
	}
}

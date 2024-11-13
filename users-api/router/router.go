package router

import (
	"users-app/users-api/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Rutas de usuarios
	r.HandleFunc("/users", controllers.RegisterUserHandler).Methods("POST")    // Registro
	r.HandleFunc("/login", controllers.LoginUserHandler).Methods("POST")       // Login
	r.HandleFunc("/users", controllers.GetAllUsersHandler).Methods("GET")      // Obtener todos los usuarios
	r.HandleFunc("/users/{id}", controllers.GetUserByIDHandler).Methods("GET") // Obtener usuario por ID
	r.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")  // Actualizar usuario

	return r
}

package router

import (
	"net/http"
	"users-app/users-api/controllers"
	"users-app/users-api/middleware"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Rutas públicas
	r.HandleFunc("/users", controllers.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUserHandler).Methods("POST")

	// Rutas protegidas
	r.Handle("/protected-route", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProtectedHandler))).Methods("GET")
	r.Handle("/admin-route", middleware.AdminMiddleware(http.HandlerFunc(controllers.AdminHandler))).Methods("GET")

	return r
}

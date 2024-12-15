package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"users-app/users-api/controllers"
	"users-app/users-api/middleware"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Rutas de usuarios
	r.HandleFunc("/users", controllers.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUserHandler).Methods("POST")
	r.HandleFunc("/users", controllers.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")

	// Ruta protegida
	r.Handle("/protected-route", middleware.AuthMiddleware(http.HandlerFunc(controllers.ProtectedHandler))).Methods("GET")

	// Ruta para listar todas las rutas disponibles
	r.HandleFunc("/routes", func(w http.ResponseWriter, r *http.Request) {
		var routes []string

		err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			path, err := route.GetPathTemplate()
			if err == nil {
				routes = append(routes, path)
			}
			return nil
		})

		if err != nil {
			http.Error(w, fmt.Sprintf("Error al recorrer las rutas: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routes)
	}).Methods("GET")

	return r
}

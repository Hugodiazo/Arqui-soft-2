// search-api/router/router.go
package router

import (
	"search-api/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/search", controllers.SearchCoursesHandler).Methods("GET")
	return router
}

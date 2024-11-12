package router

import (
	"search-app/search-api/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/search", controllers.SearchCoursesHandler).Methods("GET")
	return r
}

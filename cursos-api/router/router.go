// cursos-api/router/router.go
package router

import (
	"cursos-app/cursos-api/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define las rutas de cursos
	r.HandleFunc("/courses", controllers.GetCoursesHandler).Methods("GET")
	r.HandleFunc("/courses", controllers.CreateCourseHandler).Methods("POST")
	r.HandleFunc("/courses/{id}", controllers.GetCourseByIDHandler).Methods("GET")
	r.HandleFunc("/courses/{id}", controllers.UpdateCourseHandler).Methods("PUT")
	r.HandleFunc("/courses/enroll", controllers.EnrollCourseHandler).Methods("POST")
	r.HandleFunc("/enrollments/{user_id}", controllers.GetEnrollmentsByUserHandler).Methods("GET") // Asegúrate de que esta línea esté incluida

	return r
}

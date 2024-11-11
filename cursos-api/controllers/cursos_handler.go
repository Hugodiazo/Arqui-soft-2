// cursos-api/controllers/course_controller.go
package controllers

import (
	"cursos-app/cursos-api/domain"
	"cursos-app/cursos-api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Obtener todos los cursos
func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := services.GetAllCourses()
	if err != nil {
		http.Error(w, "Error al obtener cursos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Crear un curso
func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course domain.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	createdCourse, err := services.CreateCourse(course)
	if err != nil {
		http.Error(w, "Error al crear curso", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCourse)
}

// Obtener un curso por ID
func GetCourseByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	course, err := services.GetCourseByID(id)
	if err != nil {
		http.Error(w, "Curso no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

// Actualizar un curso por ID
func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var course domain.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	if err := services.UpdateCourse(id, course); err != nil {
		http.Error(w, "Error al actualizar el curso", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Curso actualizado con éxito"})
}

// Inscribir un usuario en un curso
func EnrollCourseHandler(w http.ResponseWriter, r *http.Request) {
	var enrollment domain.Enrollment
	if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	if err := services.EnrollCourse(enrollment); err != nil {
		http.Error(w, "Error al inscribirse en el curso", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inscripción exitosa"})
}

// Handler para obtener los cursos en los que está inscrito un usuario
func GetEnrollmentsByUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extrae el user_id desde la URL
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	// Llama a la capa de servicio para obtener los cursos inscritos
	courses, err := services.GetCoursesByUserID(userID)
	if err != nil {
		http.Error(w, "Error al obtener cursos inscritos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

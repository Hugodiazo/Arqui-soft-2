package controllers

import (
	"encoding/json"
	"net/http"
	"search-app/search-api/services"
)

// Handler para la búsqueda de cursos
func SearchCoursesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Parámetro de búsqueda 'q' es requerido", http.StatusBadRequest)
		return
	}

	courses, err := services.SearchCourses(query)
	if err != nil {
		http.Error(w, "Error al buscar cursos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

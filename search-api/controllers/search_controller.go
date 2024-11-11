// search-api/controllers/search_controller.go
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchCoursesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	searchURL := fmt.Sprintf("http://localhost:8983/solr/courses/select?q=%s&wt=json", query)

	resp, err := http.Get(searchURL)
	if err != nil {
		http.Error(w, "Error en la búsqueda de cursos", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "Error al decodificar respuesta de SolR", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

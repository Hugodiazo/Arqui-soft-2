package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"search-app/search-api/domain"
)

// SearchCourses realiza una búsqueda de cursos en Solr
func SearchCourses(query string) ([]domain.Course, error) {
	url := fmt.Sprintf("http://localhost:8983/solr/courses/select?q=title:%s&wt=json", query)
	log.Printf("Solr Query URL: %s", url)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error al buscar en Solr: %v", err)
		return nil, errors.New("error al realizar la búsqueda en Solr")
	}
	defer resp.Body.Close()

	var result struct {
		Response struct {
			Docs []domain.Course `json:"docs"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error al decodificar la respuesta de Solr: %v", err)
		return nil, errors.New("error al procesar los resultados de búsqueda")
	}

	log.Printf("Decoded Solr Response: %v", result)

	if len(result.Response.Docs) == 0 {
		return nil, errors.New("no se encontraron cursos")
	}

	return result.Response.Docs, nil
}

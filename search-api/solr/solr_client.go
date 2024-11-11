// search-api/solr/solr.go
package solr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"search-api/models"
)

const solrURL = "http://localhost:8983/solr/courses/update/json/docs?commit=true"

// Agregar o actualizar un curso en SolR
func AddOrUpdateCourse(course models.Course) error {
	body, err := json.Marshal(course)
	if err != nil {
		return fmt.Errorf("error al convertir curso a JSON: %v", err)
	}

	resp, err := http.Post(solrURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error al enviar solicitud a SolR: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("respuesta de SolR no OK: %s", resp.Status)
	}

	log.Printf("Curso %s agregado o actualizado en SolR", course.ID)
	return nil
}

// Eliminar un curso de SolR
func DeleteCourse(id string) error {
	deleteURL := fmt.Sprintf("%s/courses/update?commit=true", solrURL)
	body := fmt.Sprintf(`{"delete": {"id":"%s"}}`, id)

	resp, err := http.Post(deleteURL, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return fmt.Errorf("error al enviar solicitud de eliminación a SolR: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("respuesta de SolR no OK: %s", resp.Status)
	}

	log.Printf("Curso %s eliminado de SolR", id)
	return nil
}

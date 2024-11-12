package clients

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const SolrURL = "http://localhost:8983/solr/cursos/update?commit=true"

// IndexCourse indexa el curso en Solr
func IndexCourse(course interface{}) error {
	data, err := json.Marshal(course)
	if err != nil {
		return err
	}

	resp, err := http.Post(SolrURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("Curso indexado en Solr con éxito")
	return nil
}

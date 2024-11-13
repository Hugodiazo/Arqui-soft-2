// search-api/services/index_service.go
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"search-app/search-api/domain"
	"search-app/search-api/utils"

	"github.com/streadway/amqp"
)

// Función existente para buscar cursos en Solr
func SearchCourses(query string) ([]domain.Course, error) {
	// Escapar espacios en blanco en el query
	query = url.QueryEscape(query)

	// Construye la URL de consulta a Solr
	solrURL := fmt.Sprintf("http://localhost:8983/solr/courses/select?q=title:%s&wt=json", query)
	log.Println("Solr Query URL:", solrURL) // Imprime la URL para verificarla

	// Realiza la solicitud HTTP a Solr
	resp, err := http.Get(solrURL)
	if err != nil {
		log.Printf("Error al realizar la búsqueda en Solr: %v", err)
		return nil, fmt.Errorf("error al buscar cursos")
	}
	defer resp.Body.Close()

	// Verifica el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Error de Solr: %s", string(bodyBytes))
		return nil, fmt.Errorf("error en la respuesta de Solr: %s", resp.Status)
	}

	// Decodifica la respuesta JSON de Solr
	var result struct {
		Response struct {
			Docs []domain.Course `json:"docs"`
		} `json:"response"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error al decodificar la respuesta de Solr: %v", err)
		return nil, err
	}

	log.Printf("Cursos encontrados en Solr: %v", result.Response.Docs)
	return result.Response.Docs, nil
}

// Función para manejar mensajes de RabbitMQ
func HandleRabbitMQMessage(d amqp.Delivery) {
	log.Println("Mensaje recibido de RabbitMQ:", string(d.Body))

	var course domain.Course
	err := json.Unmarshal(d.Body, &course)
	if err != nil {
		log.Printf("Error al decodificar mensaje de RabbitMQ: %v", err)
		return
	}

	log.Println("Mensaje decodificado correctamente:", course)

	// Intenta actualizar el índice de Solr
	if err := UpdateSolrIndex(course); err != nil {
		log.Printf("Error al actualizar el índice de Solr: %v", err)
	} else {
		log.Println("Índice de Solr actualizado correctamente")
	}
}

// Función para actualizar el índice de Solr con los datos del curso
func UpdateSolrIndex(course domain.Course) error {
	url := "http://localhost:8983/solr/courses/update?commit=true"
	courseJSON, err := json.Marshal([]domain.Course{course})
	if err != nil {
		log.Printf("Error al codificar curso a JSON: %v", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(courseJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Printf("Error al crear solicitud de actualización a Solr: %v", err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error al enviar actualización a Solr: %v", err)
		return fmt.Errorf("error en la solicitud a Solr, código de respuesta: %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}

// Inicializa la conexión a RabbitMQ y la suscripción
func InitRabbitMQListener() {
	utils.ConnectRabbitMQ()
	err := utils.SubscribeToQueue("course_updates", HandleRabbitMQMessage)
	if err != nil {
		log.Fatalf("Error al suscribirse a la cola de RabbitMQ: %v", err)
	}
}

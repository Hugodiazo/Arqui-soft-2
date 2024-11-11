// search-api/services/index_service.go
package services

import (
	"encoding/json"
	"log"

	"search-api/models"
	"search-api/solr"
	"search-api/utils"

	"github.com/streadway/amqp"
)

// Procesar mensajes de RabbitMQ para actualizar el índice en SolR
func ProcessCourseUpdates() {
	utils.SubscribeToQueue("course_updates", func(d amqp.Delivery) {
		var course models.Course
		if err := json.Unmarshal(d.Body, &course); err != nil {
			log.Printf("Error al decodificar mensaje de curso: %v", err)
			return
		}

		if course.Availability {
			err := solr.AddOrUpdateCourse(course)
			if err != nil {
				log.Printf("Error al agregar/actualizar curso en SolR: %v", err)
			}
		} else {
			err := solr.DeleteCourse(course.ID)
			if err != nil {
				log.Printf("Error al eliminar curso de SolR: %v", err)
			}
		}
	})
}

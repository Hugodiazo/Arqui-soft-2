// cursos-api/clients/rabbitmq_client.go
package clients

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection

func ConnectRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("No se pudo conectar a RabbitMQ:", err)
	}

	RabbitConn = conn
	log.Println("Conectado a RabbitMQ")
}

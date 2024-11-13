// cursos-api/clients/rabbitmq_client.go
package clients

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQChannel *amqp.Channel

// ConnectRabbitMQ establece la conexión a RabbitMQ
func ConnectRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %s", err)
	}

	RabbitMQChannel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir canal en RabbitMQ: %s", err)
	}

	log.Println("Conexión a RabbitMQ establecida con éxito")
}

// PublishMessage publica un mensaje en el exchange especificado
func PublishMessage(exchange, routingKey, body string) error {
	err := RabbitMQChannel.Publish(
		exchange,   // Exchange
		routingKey, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %s", err)
		return err
	}

	log.Println("Mensaje publicado en RabbitMQ:", body)
	return nil
}

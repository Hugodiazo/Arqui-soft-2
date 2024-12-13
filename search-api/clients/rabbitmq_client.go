// clients/rabbitmq_client.go
package clients

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

// ConnectRabbitMQ establece la conexión con RabbitMQ
func ConnectRabbitMQ() {
	var err error
	RabbitMQConn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %s", err)
	}

	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir un canal en RabbitMQ: %s", err)
	}

	log.Println("Conexión a RabbitMQ establecida con éxito")
}

// PublishMessage publica un mensaje en el exchange especificado
func PublishMessage(exchange, routingKey, body string) error {
	err := RabbitMQChannel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
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

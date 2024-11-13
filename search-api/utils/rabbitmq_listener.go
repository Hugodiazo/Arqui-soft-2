package utils

import (
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func ConnectRabbitMQ() {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Error al crear canal en RabbitMQ: %v", err)
	}

	// Declarar la cola 'course_updates' para asegurarse de que exista
	_, err = ch.QueueDeclare(
		"course_updates", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola 'course_updates': %v", err)
	}
}

func CloseRabbitMQ() {
	ch.Close()
	conn.Close()
}

// Suscribirse a una cola de RabbitMQ
func SubscribeToQueue(queueName string, handler func(amqp.Delivery)) error {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(d)
		}
	}()

	return nil
}

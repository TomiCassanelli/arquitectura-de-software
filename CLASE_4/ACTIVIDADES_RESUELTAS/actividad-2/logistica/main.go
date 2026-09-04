package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const exchangeName = "ex.pedido.confirmado"
const queueName = "q.logistica"

func main() {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}
	if _, err := ch.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}

	// SOLUCIÓN DE LA ACTIVIDAD 2 — Fanout para logística
	// ================================================================

	if err := ch.QueueBind(queueName, "", exchangeName, false, nil); err != nil {
		log.Fatal(err)
	}

	// ================================================================

	messages, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[LOGÍSTICA] Lista.")
	for message := range messages {
		log.Printf("[LOGÍSTICA] Recibí %s. Preparo el envío.", message.Body)
	}
}

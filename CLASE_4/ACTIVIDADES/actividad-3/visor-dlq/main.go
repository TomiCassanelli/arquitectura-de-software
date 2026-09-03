package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const dlqName = "q.pedidos.venenosos"

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

	messages, err := ch.Consume(dlqName, "visor-dlq", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Visor de DLQ listo. CTRL+C para salir.")
	for message := range messages {
		log.Printf("[DLQ] Mensaje venenoso: %s", message.Body)
	}
}

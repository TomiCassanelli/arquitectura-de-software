package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const exchangeName = "ex.pedido.confirmado"
const queueName = "q.facturas"

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

	// TODO: declarar el exchange fanout.
	// TODO: declarar q.facturas.
	// TODO: vincular q.facturas al exchange con QueueBind.
	// Pista: la routing key de un fanout puede ser una cadena vacía.

	messages, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Facturación escuchando...")
	for message := range messages {
		log.Printf("[Facturación] Recibido: %s", message.Body)
	}
}

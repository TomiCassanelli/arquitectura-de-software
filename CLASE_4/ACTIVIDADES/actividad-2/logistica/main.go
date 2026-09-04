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

	// El mismo exchange reparte una copia del evento a cada área vinculada.
	if err := ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}
	// Logística necesita otra cola: no comparte mensajes con facturación.
	if _, err := ch.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}

	// TODO: conectar la cola de logística al exchange.
	// Usá la misma estructura que en facturación, pero con queueName = q.logistica.

	messages, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[LOGÍSTICA] Lista. Cuando la cola esté vinculada, preparará cada envío.")
	for message := range messages {
		log.Printf("[LOGÍSTICA] Recibí %s. Preparo el envío.", message.Body)
	}
}

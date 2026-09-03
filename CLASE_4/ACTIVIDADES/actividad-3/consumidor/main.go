package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const queueName = "q.pedidos"
const dlqName = "q.pedidos.venenosos"
const dlxName = "ex.pedidos.dlx"

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

	// TODO: declarar el exchange de mensajes muertos.
	// TODO: declarar la cola q.pedidos.venenosos y vincularla al exchange.
	// TODO: crear los argumentos x-dead-letter-* para q.pedidos.
	// Pista: los argumentos deben indicar el exchange y la routing key de la DLQ.

	messages, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Consumidor listo. CTRL+C para salir.")
	for message := range messages {
		var pedido map[string]interface{}
		if err := json.Unmarshal(message.Body, &pedido); err != nil {
			log.Printf("Mensaje corrupto: %s.", message.Body)
			// TODO: rechazar el mensaje sin reencolarlo.
			continue
		}
		log.Printf("Pedido válido: %v", pedido)
		// TODO: confirmar el mensaje válido.
	}
}

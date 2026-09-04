package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const queueName = "demo.pedidos"

func main() {
	log.Println("[DEMO 2] Objetivo: ver qué sucede cuando un consumidor rechaza un mensaje.")
	log.Println("[INICIO] El consumidor debe estar ejecutándose: él crea la cola principal y la DLQ.")

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
	// El consumidor crea las colas. Por eso debe iniciarse primero.
	for _, body := range []string{"pedido-1", "ERROR", "pedido-2"} {
		log.Printf("[PUBLICANDO] Envío %q a %q...", body, queueName)
		if err := ch.Publish("", queueName, false, false, amqp.Publishing{Body: []byte(body)}); err != nil {
			log.Fatal(err)
		}
		log.Printf("[ENVIADO] %q fue entregado a RabbitMQ.", body)
	}
	log.Println("[LISTO] Mirá la otra terminal: pedido-1 y pedido-2 reciben ACK; ERROR va a la DLQ.")
}

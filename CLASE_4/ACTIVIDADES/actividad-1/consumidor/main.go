package main

import (
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const queueName = "actividad-1-pedidos"

func main() {
	rand.Seed(time.Now().UnixNano())
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

	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: agregar QoS para que cada consumidor reciba un solo mensaje pendiente.
	// err = ch.Qos(1, 0, false)
	// if err != nil {
	//     log.Fatal(err)
	// }

	messages, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Consumidor listo. CTRL+C para salir.")
	for delivery := range messages {
		// Se simula que cada consumidor puede tardar un tiempo diferente.
		delay := time.Duration(rand.Intn(5)+1) * time.Second
		log.Printf("Recibido %s. Procesando durante %v", delivery.Body, delay)
		time.Sleep(delay)
		delivery.Ack(false)
		log.Println("ACK enviado")
	}
}

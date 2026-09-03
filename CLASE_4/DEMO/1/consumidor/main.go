package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672")
	if err != nil {
		log.Fatalf("Error de conexión: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error de canal: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("cola-pedidos", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al declarar cola: %v", err)
	}

	msgs, err := ch.Consume("cola-pedidos", "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al registrar consumidor: %v", err)
	}

	log.Printf("[*] Esperando pedidos. CTRL+C para salir.")

	for d := range msgs {
		log.Printf("[Consumidor] Recibido: %s", d.Body)
		time.Sleep(2 * time.Second)
	}
}

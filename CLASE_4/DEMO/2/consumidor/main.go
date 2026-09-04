package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const queueName = "demo.pedidos"
const dlqName = "demo.pedidos.dlq"

func main() {
	log.Println("[DEMO 2] Objetivo: confirmar mensajes válidos y apartar ERROR en una DLQ.")

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

	// 1. Primero declaramos la cola donde van los mensajes rechazados.
	log.Printf("[1/3] Creando o reutilizando la DLQ %q...", dlqName)
	if _, err := ch.QueueDeclare(dlqName, false, false, false, false, nil); err != nil {
		log.Fatal(err)
	}
	// 2. La cola principal indica a dónde debe enviar sus mensajes rechazados.
	log.Printf("[2/3] Creando la cola principal %q y conectándola con la DLQ...", queueName)
	args := amqp.Table{
		// "" es el exchange por defecto: usa el nombre de la cola como destino.
		"x-dead-letter-exchange": "",
		// Si se rechaza un mensaje, RabbitMQ lo publica en esta cola.
		"x-dead-letter-routing-key": dlqName,
	}
	if _, err := ch.QueueDeclare(queueName, false, false, false, false, args); err != nil {
		log.Fatal(err)
	}

	// 3. autoAck=false nos permite elegir entre ACK (aceptar) y NACK (rechazar).
	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[3/3] Escuchando %q. El siguiente log explicará cada decisión.", queueName)
	for delivery := range msgs {
		log.Printf("[RECIBIDO] Llegó %q.", delivery.Body)
		if string(delivery.Body) == "ERROR" {
			// false, false = solo este mensaje y sin reencolarlo.
			log.Println("[DECISIÓN] ERROR no se puede procesar. Envío NACK sin reencolarlo.")
			if err := delivery.Nack(false, false); err != nil {
				log.Printf("[ERROR] No se pudo enviar el NACK: %v", err)
				continue
			}
			log.Printf("[DLQ] RabbitMQ apartó ERROR en %q. La cola principal puede continuar.", dlqName)
			continue
		}
		log.Printf("[DECISIÓN] %q es válido. Lo proceso y envío ACK.", delivery.Body)
		if err := delivery.Ack(false); err != nil {
			log.Printf("[ERROR] No se pudo enviar el ACK: %v", err)
			continue
		}
		log.Println("[ACK] Trabajo terminado. Sigo esperando el próximo mensaje.")
	}
}

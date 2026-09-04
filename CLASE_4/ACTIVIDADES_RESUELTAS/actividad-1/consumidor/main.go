package main

import (
	"log"
	"os"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const queueName = "actividad-1-pedidos"

func main() {
	consumerName := os.Getenv("CONSUMIDOR")
	if consumerName == "" {
		consumerName = "sin-nombre"
	}
	processingSeconds, err := strconv.Atoi(os.Getenv("DELAY_SEGUNDOS"))
	if err != nil || processingSeconds < 1 {
		processingSeconds = 2
	}

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

	if _, err = ch.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}

	// SOLUCIÓN DE LA ACTIVIDAD 1 — Fair Dispatch con QoS
	// ================================================================

	if err = ch.Qos(1, 0, false); err != nil {
		log.Fatal(err)
	}

	// ================================================================

	messages, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[%s] Listo. Tarda %ds por pedido. CTRL+C para salir.", consumerName, processingSeconds)
	for delivery := range messages {
		delay := time.Duration(processingSeconds) * time.Second
		log.Printf("[%s] Recibido %s. Procesando durante %v", consumerName, delivery.Body, delay)
		time.Sleep(delay)
		if err := delivery.Ack(false); err != nil {
			log.Printf("[%s] Error al enviar ACK: %v", consumerName, err)
			continue
		}
		log.Printf("[%s] ACK enviado", consumerName)
	}
}

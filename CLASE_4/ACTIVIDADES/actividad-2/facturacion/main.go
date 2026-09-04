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

	// El exchange es el punto común donde la tienda publica los pedidos confirmados.
	if err := ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}
	// Facturación tiene su propia cola: así puede trabajar a su ritmo.
	if _, err := ch.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}

	// TODO: conectar la cola de facturación al exchange.
	// Pista: QueueBind(queueName, "", exchangeName, false, nil)
	// En fanout, "" está bien: el exchange envía una copia a todas las colas vinculadas.

	messages, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[FACTURACIÓN] Lista. Cuando la cola esté vinculada, recibirá cada pedido confirmado.")
	for message := range messages {
		log.Printf("[FACTURACIÓN] Recibí %s. Genero la factura.", message.Body)
	}
}

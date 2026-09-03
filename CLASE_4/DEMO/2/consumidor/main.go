package main

import (
	"encoding/json"
	"log"

	"rabbitmq-demo/shared"
)

func main() {
	conn, err := shared.Connect(shared.AMQPURI)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	q, err := shared.DeclareQueues(ch)
	if err != nil {
		log.Fatal(err)
	}

	// autoAck=false permite decidir si el mensaje se confirma o va a la DLQ.
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[Consumidor] Escuchando la cola '%s'.", q.Name)
	for delivery := range msgs {
		var pedido shared.Pedido
		if err := json.Unmarshal(delivery.Body, &pedido); err != nil {
			log.Printf("[Consumidor] Error al leer el pedido: %v", err)
			// NACK con requeue=false: RabbitMQ lo desvía automáticamente a la DLQ.
			delivery.Nack(false, false)
			continue
		}
		log.Printf("[Consumidor] Recibido: ID=%s, monto=%.2f, cliente=%s", pedido.PedidoID, pedido.Monto, pedido.Cliente)
		delivery.Ack(false)
	}
}

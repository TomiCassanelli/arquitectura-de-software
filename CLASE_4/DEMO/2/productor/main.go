package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"rabbitmq-demo/shared"

	amqp "github.com/rabbitmq/amqp091-go"
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

	// Se envía un mensaje correcto para mostrar el flujo normal.
	pedido := shared.Pedido{PedidoID: "PED-1", Monto: 150.50, Cliente: "C-123"}
	body, err := json.Marshal(pedido)
	if err != nil {
		log.Fatal(err)
	}
	publish(ch, q.Name, body)

	// Se envía un mensaje incorrecto para probar la DLQ.
	publish(ch, q.Name, []byte("esto no es JSON"))
}

func publish(ch *amqp.Channel, queueName string, body []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ch.PublishWithContext(ctx, "", queueName, false, false, amqp.Publishing{ContentType: "application/json", Body: body}); err != nil {
		log.Fatal(err)
	}
	log.Printf("[Productor] Publicado: %s", body)
}

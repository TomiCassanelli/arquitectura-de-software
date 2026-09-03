package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const exchangeName = "ex.pedido.confirmado"

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

	// TODO: declarar un exchange durable de tipo fanout antes de publicar.
	// Pista: ExchangeDeclare(exchangeName, "fanout", ...)

	body := []byte(`{"pedido_id":"PED-1","estado":"confirmado"}`)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx, exchangeName, "", false, false, amqp.Publishing{ContentType: "application/json", Body: body})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Publicado en fanout: %s", body)
}

package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const queueName = "q.pedidos"

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
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	messages := [][]byte{
		[]byte(`{"pedido_id":"PED-1","monto":100}`),
		[]byte(`esto no es JSON`),
		[]byte(`{"pedido_id":"PED-2","monto":250}`),
	}
	for _, body := range messages {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = ch.PublishWithContext(ctx, "", queueName, false, false, amqp.Publishing{ContentType: "application/json", Body: body})
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Publicado: %s", body)
		time.Sleep(1 * time.Second)
	}
}

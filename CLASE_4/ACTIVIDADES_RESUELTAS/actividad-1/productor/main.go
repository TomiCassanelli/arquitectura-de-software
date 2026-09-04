package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const queueName = "actividad-1-pedidos"

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

	if _, err = ch.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}

	for number := 1; number <= 10; number++ {
		body := fmt.Sprintf("pedido-%d", number)
		err = ch.PublishWithContext(context.Background(), "", queueName, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			Body:         []byte(body),
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Publicado: %s", body)
		time.Sleep(time.Second)
	}
}

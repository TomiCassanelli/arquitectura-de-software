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

	// El productor también declara el exchange para poder ejecutarse por separado.
	// No es parte del ejercicio: el foco está en los vínculos de cada consumidor.
	if err := ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil); err != nil {
		log.Fatal(err)
	}

	body := []byte(`{"pedido_id":"PED-1","estado":"confirmado"}`)
	log.Printf("[TIENDA] El pedido fue confirmado: %s", body)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx, exchangeName, "", false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[TIENDA] Evento publicado en el exchange fanout.")
	log.Println("[OBSERVÁ] Cada cola vinculada recibirá su propia copia del pedido.")
}

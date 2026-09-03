package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672")
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir canal: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("cola-pedidos", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al declarar cola: %v", err)
	}

	body := `{"pedido_id": "PED-10293", "monto": 150.50, "cliente": "C-123"}`

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "application/json", Body: []byte(body)})
	if err != nil {
		log.Fatalf("Error al publicar mensaje: %v", err)
	}
	log.Printf("[Productor] Evento publicado: %s", body)
}

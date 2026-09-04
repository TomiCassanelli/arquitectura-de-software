package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const queueName = "cola-pedidos"

func main() {
	log.Println("[DEMO 1] Objetivo: publicar un pedido y seguir el mensaje en RabbitMQ.")

	// 1. Nos conectamos al broker (RabbitMQ).
	log.Println("[1/4] Conectando con RabbitMQ...")
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672")
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer conn.Close()

	// 2. Abrimos un canal: es por donde enviamos las operaciones a RabbitMQ.
	log.Println("[2/4] Conexión lograda. Abriendo un canal...")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir canal: %v", err)
	}
	defer ch.Close()

	// 3. Creamos la cola si todavía no existe.
	// Los últimos parámetros son detalles de configuración; no cambian en esta demo.
	log.Printf("[3/4] Verificando la cola %q...", queueName)
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al declarar cola: %v", err)
	}

	body := "pedido-101"

	// 4. Publicamos el texto en la cola. "" significa exchange por defecto.
	log.Printf("[4/4] Publicando %q en la cola...", body)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{Body: []byte(body)})
	if err != nil {
		log.Fatalf("Error al publicar mensaje: %v", err)
	}
	log.Printf("[LISTO] %q fue enviado. Si el consumidor no está activo, mirá Ready=1 en el panel.", body)
}

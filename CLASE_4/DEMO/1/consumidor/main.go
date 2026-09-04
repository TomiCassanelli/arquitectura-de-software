package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const queueName = "cola-pedidos"

func main() {
	log.Println("[DEMO 1] Objetivo: recibir un pedido y confirmar cuándo termina el trabajo.")

	// 1. Conexión y canal: igual que en el productor.
	log.Println("[1/3] Conectando con RabbitMQ...")
	conn, err := amqp.Dial("amqp://user:pass@localhost:5672")
	if err != nil {
		log.Fatalf("Error de conexión: %v", err)
	}
	defer conn.Close()

	log.Println("[1/3] Conexión lograda. Abriendo un canal...")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error de canal: %v", err)
	}
	defer ch.Close()

	// 2. El consumidor también declara la cola. Si ya existe, RabbitMQ no la duplica.
	// Los últimos parámetros son detalles de configuración; no cambian en esta demo.
	log.Printf("[2/3] Verificando la cola %q...", queueName)
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al declarar cola: %v", err)
	}

	// autoAck=false: el consumidor confirma recién cuando termina de procesar.
	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error al registrar consumidor: %v", err)
	}

	log.Printf("[3/3] Esperando pedidos. En el panel, Ready bajará y Unacked subirá mientras proceso. CTRL+C para salir.")

	for d := range msgs {
		// 3. Recibir no es terminar: simulamos el trabajo antes de confirmar.
		log.Printf("[RECIBIDO] %q. Ahora está Unacked: el trabajo todavía no terminó.", d.Body)
		log.Println("[PROCESANDO] Simulando 10 segundos de trabajo...")
		time.Sleep(10 * time.Second)
		if err := d.Ack(false); err != nil {
			log.Printf("[Consumidor] Error al enviar ACK: %v", err)
			continue
		}
		log.Println("[ACK] Trabajo terminado. El mensaje ya no figura como pendiente en la cola.")
	}
}

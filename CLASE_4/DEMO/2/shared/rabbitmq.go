package shared

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(uri string) (*amqp.Connection, error) {
	for attempt := 1; attempt <= MaxConnections; attempt++ {
		log.Printf("[Conexión] Intento %d de %d", attempt, MaxConnections)
		conn, err := amqp.Dial(uri)
		if err == nil {
			log.Println("[Conexión] Conectado a RabbitMQ")
			return conn, nil
		}

		if attempt < MaxConnections {
			delay := time.Duration(attempt*BaseDelayMs) * time.Millisecond
			log.Printf("[Conexión] Error: %v. Reintentando en %v", err, delay)
			time.Sleep(delay)
		}
	}

	return nil, fmt.Errorf("no se pudo conectar después de %d intentos", MaxConnections)
}

func DeclareQueues(ch *amqp.Channel) (amqp.Queue, error) {
	// La DLQ guarda los mensajes que el consumidor rechaza sin reencolarlos.
	if _, err := ch.QueueDeclare(DLQName, true, false, false, false, nil); err != nil {
		return amqp.Queue{}, err
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": DLQName,
	}
	return ch.QueueDeclare(QueueName, true, false, false, false, args)
}

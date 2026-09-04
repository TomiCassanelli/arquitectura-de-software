package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const amqpURI = "amqp://user:pass@localhost:5672"
const maxAttempts = 4

func main() {
	log.Println("[DEMO 3] Objetivo: reintentar una conexión sin saturar RabbitMQ.")
	log.Println("[PREPARACIÓN] Para observar los fallos, iniciá esta demo con RabbitMQ apagado.")
	conn, err := connectWithBackoff(amqpURI)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("[ÉXITO] RabbitMQ respondió. La demo terminó: solo mostramos la conexión inicial.")
}

// connectWithBackoff intenta conectar más de una vez.
// Si falla, espera 1 s, luego 2 s y luego 4 s antes de volver a intentar.
func connectWithBackoff(uri string) (*amqp.Connection, error) {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		log.Printf("[INTENTO %d/%d] Intento conectar con RabbitMQ...", attempt, maxAttempts)

		// Intentamos la misma conexión normal que usamos en las otras demos.
		conn, err := amqp.Dial(uri)
		if err == nil {
			return conn, nil
		}

		if attempt == maxAttempts {
			return nil, fmt.Errorf("RabbitMQ sigue sin responder después de %d intentos", maxAttempts)
		}

		// 1<<(attempt-1) produce 1, 2 y 4. Eso es backoff exponencial.
		wait := time.Second * time.Duration(1<<(attempt-1))
		log.Printf("[BACKOFF] RabbitMQ todavía no responde. Espero %v antes de reintentar.", wait)
		time.Sleep(wait)
	}

	return nil, fmt.Errorf("no se pudo conectar")
}

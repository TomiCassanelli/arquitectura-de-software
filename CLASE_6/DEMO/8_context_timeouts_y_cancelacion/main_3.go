package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Las tres consultas comparten el mismo límite de tiempo.
	requestContext, cancelRequest := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancelRequest()

	resultChannel := make(chan string)

	go fetchData(requestContext, "Price", 100*time.Millisecond, resultChannel)
	go fetchData(requestContext, "Stock", 200*time.Millisecond, resultChannel)
	go fetchData(requestContext, "Reviews", 600*time.Millisecond, resultChannel)

	// Cada consulta envía un resultado: recibido o cancelado.
	for receivedResults := 0; receivedResults < 3; receivedResults++ {
		fmt.Println(<-resultChannel)
	}
}

func fetchData(requestContext context.Context, dataName string, delay time.Duration, resultChannel chan string) {
	select {
	case <-time.After(delay):
		// Esta fuente alcanzó a responder antes del timeout compartido.
		resultChannel <- dataName + ": received"
	case <-requestContext.Done():
		// Esta fuente recibió el mismo aviso de cancelación.
		resultChannel <- dataName + ": canceled"
	}
}

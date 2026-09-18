package main

import (
	"fmt"
	"time"
)

func main() {
	// Ahora el channel tiene espacio para dos pedidos.
	orderChannel := make(chan string, 2)

	go sendOrderTo(orderChannel)

	time.Sleep(5 * time.Second)
	fmt.Println("Receiver: now I am ready")

	receivedOrder := <-orderChannel
	fmt.Println("Receiver:", receivedOrder)
}

func sendOrderTo(orderChannel chan string) {
	fmt.Println("Sender: I want to send the order")
	orderChannel <- "Order is ready"
	// Con buffer, este mensaje aparece antes de que main reciba.
	fmt.Println("Sender: the order is waiting in the channel")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Volvemos a un channel sin buffer para aplicar el envío/recepción a varios pedidos.
	orderChannel := make(chan string)

	go sendOrdersTo(orderChannel)
	go receiveOrdersFrom(orderChannel)

	// Solo evita que main termine antes de que se vean los dos pedidos.
	time.Sleep(2 * time.Second)
}

func sendOrdersTo(orderChannel chan string) {
	for orderID := 1; orderID <= 2; orderID++ {
		fmt.Println("Sender: I want to send order", orderID)
		orderChannel <- fmt.Sprintf("Order %d is ready", orderID)
		fmt.Println("Sender: order", orderID, "was received")
	}
}

func receiveOrdersFrom(orderChannel chan string) {
	for receivedOrders := 0; receivedOrders < 2; receivedOrders++ {
		receivedOrder := <-orderChannel
		fmt.Println("Receiver:", receivedOrder)
	}
}

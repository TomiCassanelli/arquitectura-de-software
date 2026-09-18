package main

import (
	"fmt"
	"time"
)

func main() {
	// Ahora hay una fila para dos pedidos.
	orderChannel := make(chan string, 2)

	go sendOrdersTo(orderChannel)
	go receiveOrdersFrom(orderChannel)

	time.Sleep(4 * time.Second)
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

		time.Sleep(time.Second)
	}
}

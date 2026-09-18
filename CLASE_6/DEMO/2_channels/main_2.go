package main

import (
	"fmt"
	"time"
)

func main() {
	orderChannel := make(chan string)

	go sendOrderTo(orderChannel)

	// main todavía no recibe: el envío queda esperando aquí cinco segundos.
	time.Sleep(5 * time.Second)
	fmt.Println("Receiver: now I am ready")

	receivedOrder := <-orderChannel
	fmt.Println("Receiver:", receivedOrder)
}

func sendOrderTo(orderChannel chan string) {
	fmt.Println("Sender: I want to send the order")
	orderChannel <- "Order is ready"
	fmt.Println("Sender: the order was received, I can continue")
}

package main

import "fmt"

func main() {
	// Este channel transporta textos sobre un pedido.
	orderChannel := make(chan string)

	go sendOrderTo(orderChannel)

	// <- significa: recibir lo que llegue por el channel.
	receivedOrder := <-orderChannel
	fmt.Println("Received:", receivedOrder)
}

func sendOrderTo(orderChannel chan string) {
	// <- significa: enviar este texto hacia el channel.
	orderChannel <- "Order is ready"
}

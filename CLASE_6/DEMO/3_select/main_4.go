package main

import "fmt"

func main() {
	// No hay nadie que envíe mensajes a este channel.
	messageChannel := make(chan string)

	// Sin default, select espera para siempre un mensaje que no llegará.
	// Al no quedar ninguna goroutine que pueda avanzar, Go muestra:
	// fatal error: all goroutines are asleep - deadlock!
	select {
	case message := <-messageChannel:
		fmt.Println("Received:", message)
	}
}

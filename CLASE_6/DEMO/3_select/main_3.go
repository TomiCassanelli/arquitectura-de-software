package main

import (
	"fmt"
	"time"
)

func main() {
	messageChannel := make(chan string)
	exitChannel := make(chan bool)

	go sendMessages(messageChannel)
	go sendExit(exitChannel)

	// El for vuelve a ejecutar select después de cada mensaje.
	for {
		select {
		case message := <-messageChannel:
			fmt.Println("Message received:", message)
		case exitSignal := <-exitChannel:
			fmt.Println("Select finished:", exitSignal)
			return
		}
	}
}

func sendMessages(messageChannel chan string) {
	// La misma funcion, pero ahora envía varios mensajes.
	for messageID := 0; ; messageID++ {
		messageChannel <- fmt.Sprintf("Hello %d", messageID)
		time.Sleep(500 * time.Millisecond)
	}
}

func sendExit(exitChannel chan bool) {
	time.Sleep(5 * time.Second)
	exitChannel <- true
}

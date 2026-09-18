package main

import (
	"fmt"
	"time"
)

func main() {
	messageChannel := make(chan string)
	exitChannel := make(chan bool)

	go sendMessage(messageChannel)
	go sendExit(exitChannel)

	// Solo cambiaron las demoras, ahora exit llegará primero.
	select {
	case message := <-messageChannel:
		fmt.Println("Message received:", message)
	case exitSignal := <-exitChannel:
		fmt.Println("Exit received:", exitSignal)
	}
}

func sendMessage(messageChannel chan string) {
	time.Sleep(2 * time.Second)
	messageChannel <- "Hello"
}

func sendExit(exitChannel chan bool) {
	time.Sleep(time.Second)
	exitChannel <- true
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Select escuchará estos dos channels.
	messageChannel := make(chan string)
	exitChannel := make(chan bool)

	go sendMessage(messageChannel)
	go sendExit(exitChannel)

	// El primer channel que tenga algo listo gana.
	select {
	case message := <-messageChannel:
		fmt.Println("Message received:", message)
	case exitSignal := <-exitChannel:
		fmt.Println("Exit received:", exitSignal)
	}
}

func sendMessage(messageChannel chan string) {
	time.Sleep(time.Second)
	messageChannel <- "Hello"
}

func sendExit(exitChannel chan bool) {
	time.Sleep(2 * time.Second)
	exitChannel <- true
}

package main

import (
	"fmt"
	"time"
)

func main() {
	serviceChannel := make(chan string)

	go sendFastThenSlowResponse(serviceChannel)

	responseTimeout := time.After(time.Second)

	for {
		select {
		case response := <-serviceChannel:
			fmt.Println("Response:", response)

		case <-responseTimeout:
			fmt.Println("Timeout")
			return
		}
	}
}

func sendFastThenSlowResponse(channel chan string) {
	// La primera llega a tiempo.
	time.Sleep(500 * time.Millisecond)
	channel <- "First response"

	// La segunda llega demasiado tarde para el mismo timeout.
	time.Sleep(2 * time.Second)
	channel <- "Second response"
}

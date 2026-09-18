package main

import (
	"fmt"
	"time"
)

func main() {
	serviceChannel := make(chan string)

	go sendFastResponse(serviceChannel)

	responseTimeout := time.After(time.Second)

	select {
	case response := <-serviceChannel:
		fmt.Println("Response:", response)
	case <-responseTimeout:
		fmt.Println("Timeout")
	}
}

func sendFastResponse(channel chan string) {
	// Esta respuesta llega antes del límite de un segundo.
	time.Sleep(500 * time.Millisecond)
	channel <- "Service responded"
}

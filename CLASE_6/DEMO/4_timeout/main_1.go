package main

import (
	"fmt"
	"time"
)

func main() {
	// Este channel recibirá la respuesta del servicio.
	serviceChannel := make(chan string)

	go sendSlowResponse(serviceChannel)

	// Este channel se activa solo después de un segundo.
	responseTimeout := time.After(time.Second)

	select {
	case response := <-serviceChannel:
		fmt.Println("Response:", response)
	case <-responseTimeout:
		fmt.Println("Timeout")
	}
}

func sendSlowResponse(channel chan string) {
	// El servicio tarda más que el tiempo permitido.
	time.Sleep(2 * time.Second)
	channel <- "Service responded"
}

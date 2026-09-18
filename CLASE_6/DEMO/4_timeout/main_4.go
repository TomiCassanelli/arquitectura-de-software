package main

import (
	"fmt"
	"time"
)

func main() {
	serviceChannel := make(chan string)

	go sendResponses(serviceChannel)

	programTimeout := time.After(5 * time.Second)

	for {
		select {
		case response := <-serviceChannel:
			fmt.Println("Response:", response)

		case <-programTimeout:
			fmt.Println("Program timeout")
			return
		}
	}
}

func sendResponses(channel chan string) {
	// El servicio sigue enviando respuestas hasta que main termine el programa.
	for responseID := 0; ; responseID++ {
		time.Sleep(700 * time.Millisecond)
		channel <- fmt.Sprintf("Response %d", responseID)
	}
}

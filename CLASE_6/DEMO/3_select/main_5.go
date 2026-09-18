package main

import "fmt"

func main() {
	messageChannel := make(chan string)

	// Default permite continuar sin esperar.
	select {
	case message := <-messageChannel:
		fmt.Println("Received:", message)
	default:
		fmt.Println("No message available: select did not block")
	}
}

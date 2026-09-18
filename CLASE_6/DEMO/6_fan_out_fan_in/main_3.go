package main

import (
	"fmt"
	"time"
)

func main() {
	// Un solo channel recibe las respuestas de los dos comercios.
	pricesChannel := make(chan string)

	// Fan-out: enviamos la misma tarea de buscar precio a dos comercios.
	go searchStore("Frávega", 1200*time.Millisecond, pricesChannel)
	go searchStore("Garbarino", 500*time.Millisecond, pricesChannel)

	// Fan-in: las dos respuestas vuelven al mismo lugar.
	for receivedPrices := 0; receivedPrices < 2; receivedPrices++ {
		fmt.Println(<-pricesChannel)
	}
}

func searchStore(store string, delay time.Duration, priceChannel chan string) {
	time.Sleep(delay)
	priceChannel <- fmt.Sprintf("Price from %s: $1200", store)
}

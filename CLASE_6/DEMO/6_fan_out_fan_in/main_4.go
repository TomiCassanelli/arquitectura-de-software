package main

import (
	"fmt"
	"time"
)

func main() {
	pricesChannel := make(chan string)

	// Fan-out: una búsqueda sale hacia tres comercios al mismo tiempo.
	go searchStore("Frávega", 1200*time.Millisecond, pricesChannel)
	go searchStore("Garbarino", 500*time.Millisecond, pricesChannel)
	go searchStore("Musimundo", 800*time.Millisecond, pricesChannel)

	// Fan-in: main junta las tres respuestas en pricesChannel.
	for receivedPrices := 0; receivedPrices < 3; receivedPrices++ {
		fmt.Println(<-pricesChannel)
	}
}

func searchStore(store string, delay time.Duration, priceChannel chan string) {
	time.Sleep(delay)
	priceChannel <- fmt.Sprintf("Price from %s: $1200", store)
}

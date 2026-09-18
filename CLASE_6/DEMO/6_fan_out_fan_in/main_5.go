package main

import (
	"fmt"
	"time"
)

func main() {
	pricesChannel := make(chan string)

	// El mismo patrón puede consultar más comercios.
	go searchStore("Frávega", 1200*time.Millisecond, pricesChannel)
	go searchStore("Garbarino", 500*time.Millisecond, pricesChannel)
	go searchStore("Musimundo", 800*time.Millisecond, pricesChannel)
	go searchStore("Coto", 1500*time.Millisecond, pricesChannel)
	go searchStore("Carrefour", 1000*time.Millisecond, pricesChannel)

	for receivedPrices := 0; receivedPrices < 5; receivedPrices++ {
		fmt.Println(<-pricesChannel)
	}
}

func searchStore(store string, delay time.Duration, priceChannel chan string) {
	time.Sleep(delay)
	priceChannel <- fmt.Sprintf("Price from %s: $1200", store)
}

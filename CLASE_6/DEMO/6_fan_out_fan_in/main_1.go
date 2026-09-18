package main

import (
	"fmt"
	"time"
)

func main() {
	// El channel se crea acá, donde vamos a recibir el precio.
	priceChannel := make(chan string)

	go searchStore("Frávega", 1200*time.Millisecond, priceChannel)

	price := <-priceChannel
	fmt.Println(price)
}

func searchStore(store string, delay time.Duration, priceChannel chan string) {
	time.Sleep(delay)
	priceChannel <- fmt.Sprintf("Price from %s: $1200", store)
}

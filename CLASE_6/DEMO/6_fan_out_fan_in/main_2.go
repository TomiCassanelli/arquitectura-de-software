package main

import (
	"fmt"
	"time"
)

func main() {
	// Cada comercio tiene su propio channel, creado en main.
	fravegaChannel := make(chan string)
	garbarinoChannel := make(chan string)

	// Las dos consultas empiezan a la vez: este es el comienzo del fan-out.
	go searchStore("Frávega", 1200*time.Millisecond, fravegaChannel)
	go searchStore("Garbarino", 500*time.Millisecond, garbarinoChannel)

	// main lee primero Frávega y luego Garbarino, siempre en ese orden.
	fmt.Println(<-fravegaChannel)
	fmt.Println(<-garbarinoChannel)
}

func searchStore(store string, delay time.Duration, priceChannel chan string) {
	time.Sleep(delay)
	priceChannel <- fmt.Sprintf("Price from %s: $1200", store)
}

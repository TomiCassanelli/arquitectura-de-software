package main

import (
	"fmt"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	firstStoreName  = "XXX"
	firstStorePrice = XXX
	firstStoreDelay = XXX * time.Millisecond

	secondStoreName  = "XXX"
	secondStorePrice = XXX
	secondStoreDelay = XXX * time.Millisecond
)

type Price struct {
	Store string
	Value int
}

func searchFirstStore() Price {
	time.Sleep(firstStoreDelay)
	return Price{Store: firstStoreName, Value: firstStorePrice}
}

func searchSecondStore() Price {
	time.Sleep(secondStoreDelay)
	return Price{Store: secondStoreName, Value: secondStorePrice}
}

// TODO: iniciar las dos búsquedas en goroutines y reunir ambas respuestas.
func comparePrices() []Price {
	return []Price{searchFirstStore(), searchSecondStore()}
}

func main() {
	startedAt := time.Now()
	for _, price := range comparePrices() {
		fmt.Printf("%s: $%d\n", price.Store, price.Value)
	}
	fmt.Println("Time:", time.Since(startedAt))
}

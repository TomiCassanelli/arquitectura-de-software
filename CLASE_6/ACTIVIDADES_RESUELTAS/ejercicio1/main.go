package main

import (
	"fmt"
	"time"
)

const (
	firstStoreName   = "Frávega"
	firstStorePrice  = 1200
	firstStoreDelay  = 800 * time.Millisecond
	secondStoreName  = "Garbarino"
	secondStorePrice = 1150
	secondStoreDelay = 400 * time.Millisecond
)

type Price struct {
	Store string
	Value int
}

func searchFirstStore() Price {
	time.Sleep(firstStoreDelay)
	return Price{firstStoreName, firstStorePrice}
}
func searchSecondStore() Price {
	time.Sleep(secondStoreDelay)
	return Price{secondStoreName, secondStorePrice}
}

func comparePrices() []Price {
	priceChannel := make(chan Price, 2)
	go sendFirstPrice(priceChannel)
	go sendSecondPrice(priceChannel)
	return []Price{<-priceChannel, <-priceChannel}
}

func sendFirstPrice(priceChannel chan Price)  { priceChannel <- searchFirstStore() }
func sendSecondPrice(priceChannel chan Price) { priceChannel <- searchSecondStore() }

func main() {
	startedAt := time.Now()
	for _, price := range comparePrices() {
		fmt.Printf("%s: $%d\n", price.Store, price.Value)
	}
	fmt.Println("Time:", time.Since(startedAt))
}

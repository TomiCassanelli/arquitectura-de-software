package main

import (
	"fmt"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	productName    = "XXX"
	reviewsScore   = XXX
	reviewsDelay   = XXX * time.Millisecond
	timeoutLimit   = XXX * time.Millisecond
	defaultReviews = XXX
)

func getReviews() int {
	time.Sleep(reviewsDelay)
	return reviewsScore
}

// TODO: esperar reviews o el timeout. Si gana el timeout, devolver defaultReviews.
func getProductReviews() int {
	return getReviews()
}

func main() {
	startedAt := time.Now()
	reviews := getProductReviews()
	fmt.Printf("%s reviews: %d\n", productName, reviews)
	fmt.Println("Time:", time.Since(startedAt))
}

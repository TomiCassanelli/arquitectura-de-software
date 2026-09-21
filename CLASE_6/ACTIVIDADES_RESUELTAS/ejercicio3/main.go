package main

import (
	"context"
	"fmt"
	"time"
)

const (
	productName    = "Notebook"
	reviewsScore   = 45
	reviewsDelay   = 2 * time.Second
	timeoutLimit   = 500 * time.Millisecond
	defaultReviews = 0
)

func getReviews() int { time.Sleep(reviewsDelay); return reviewsScore }

func getProductReviews() int {
	reviewsChannel := make(chan int, 1)
	go sendReviews(reviewsChannel)
	requestContext, cancelRequest := context.WithTimeout(context.Background(), timeoutLimit)
	defer cancelRequest()

	select {
	case reviews := <-reviewsChannel:
		return reviews
	case <-requestContext.Done():
		return defaultReviews
	}
}

func sendReviews(reviewsChannel chan int) { reviewsChannel <- getReviews() }

func main() {
	startedAt := time.Now()
	reviews := getProductReviews()
	fmt.Printf("%s reviews: %d\n", productName, reviews)
	fmt.Println("Time:", time.Since(startedAt))
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Esta vez el context se cancela solo a los 300 ms.
	requestContext, cancelRequest := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancelRequest()

	finishedChannel := make(chan bool)

	go fetchReviews(requestContext, finishedChannel)

	// main no cancela nada: espera a que el timeout avise a Reviews.
	<-finishedChannel
}

func fetchReviews(requestContext context.Context, finishedChannel chan bool) {
	fmt.Println("Reviews: fetching")

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Reviews: received")
	case <-requestContext.Done():
		fmt.Println("Reviews: canceled")
	}

	finishedChannel <- true
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Este context se cancela cuando nosotros llamamos a cancelRequest().
	requestContext, cancelRequest := context.WithCancel(context.Background())

	// Este channel no lleva reviews: solo confirma que la tarea terminó.
	finishedChannel := make(chan bool)

	go fetchReviews(requestContext, finishedChannel)

	// La persona cerró la pantalla antes de que llegaran las reviews.
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Cancelo el contexto de manera manual")
	cancelRequest()

	// Esperamos a que la consulta reciba el aviso y termine.
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

	// Avisamos a main que ya atendimos la cancelación y terminamos.
	finishedChannel <- true
}

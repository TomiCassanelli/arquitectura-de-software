package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerCount  = 3
	processDelay = 300 * time.Millisecond
)

var productIDs = []int{101, 102, 103, 104, 105}

func updateProduct(workerID int, productID int) {
	fmt.Printf("Worker %d updated product %d\n", workerID, productID)
	time.Sleep(processDelay)
}

func processWithWorkers() {
	productChannel := make(chan int)
	var finishedWorkers sync.WaitGroup
	for workerID := 1; workerID <= workerCount; workerID++ {
		finishedWorkers.Add(1)
		go updateProducts(workerID, productChannel, &finishedWorkers)
	}
	for _, productID := range productIDs {
		productChannel <- productID
	}
	close(productChannel)
	finishedWorkers.Wait()
}

func updateProducts(workerID int, productChannel chan int, finishedWorkers *sync.WaitGroup) {
	defer finishedWorkers.Done()
	for productID := range productChannel {
		updateProduct(workerID, productID)
	}
}

func main() {
	startedAt := time.Now()
	processWithWorkers()
	fmt.Println("Time:", time.Since(startedAt))
}

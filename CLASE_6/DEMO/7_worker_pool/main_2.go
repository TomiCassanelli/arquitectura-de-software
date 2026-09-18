package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	productIDs := []int{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113}
	var finishedUpdates sync.WaitGroup

	// Hay una tarea pendiente por cada producto.
	finishedUpdates.Add(len(productIDs))

	for _, productID := range productIDs {
		// Cada producto tiene su propia goroutine.
		go updateProductPriceAndNotify(productID, &finishedUpdates)
	}

	// main espera que terminen todas las actualizaciones.
	finishedUpdates.Wait()
	fmt.Println("All products updated")
}

func updateProductPriceAndNotify(productID int, finishedUpdates *sync.WaitGroup) {
	defer finishedUpdates.Done()

	fmt.Println("Updating product", productID)
	time.Sleep(500 * time.Millisecond)
}

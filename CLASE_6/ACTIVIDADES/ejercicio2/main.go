package main

import (
	"fmt"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	workerCount  = XXX
	processDelay = XXX * time.Millisecond
)

var productIDs = []int{XXX}

func updateProduct(workerID int, productID int) {
	fmt.Printf("Worker %d updated product %d\n", workerID, productID)
	time.Sleep(processDelay)
}

// TODO: crear un channel de productos y workerCount workers fijos.
func processWithWorkers() {
	for _, productID := range productIDs {
		updateProduct(1, productID)
	}
}

func main() {
	startedAt := time.Now()
	processWithWorkers()
	fmt.Println("Time:", time.Since(startedAt))
}

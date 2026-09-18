package main

import (
	"fmt"
	"sync"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	initialStock = XXX
	totalSales   = XXX
	saleDelay    = XXX * time.Millisecond
)

// Simula totalSales compras concurrentes del mismo producto, decrementando un
// contador de stock compartido SIN ninguna protección. Corré este archivo
// con -race para ver el reporte de data race:
//
//	go run -race main.go
func main() {
	stock := initialStock

	var wg sync.WaitGroup
	wg.Add(totalSales)
	for i := 0; i < totalSales; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(saleDelay) // simula el tiempo de la "compra"
			stock--               // ← data race: compras sin sincronizar
		}()
	}
	wg.Wait()

	fmt.Println("Stock final:", stock) // valor impredecible, puede no ser 50
}

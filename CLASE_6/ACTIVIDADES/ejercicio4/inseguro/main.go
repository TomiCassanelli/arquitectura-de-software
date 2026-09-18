package main

import (
	"fmt"
	"sync"
	"time"
)

// Simula 50 compras concurrentes del mismo producto, decrementando un
// contador de stock compartido SIN ninguna protección. Corré este archivo
// con -race para ver el reporte de data race:
//
//	go run -race main.go
func main() {
	stock := 100

	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond) // simula el tiempo de la "compra"
			stock--                      // ← data race: 50 goroutines sin sincronizar
		}()
	}
	wg.Wait()

	fmt.Println("Stock final:", stock) // valor impredecible, puede no ser 50
}

package main

import (
	"fmt"
	"sync"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	initialStock = XXX
	totalSales   = XXX
)

// Hands On 2: corregir la race condition del Hands On 1 con sync.Mutex.
//
// Mismo escenario que handson1/race.go, pero protegiendo el acceso a stock
// con un Mutex. Corré:
//
//	go run -race race_mutex.go
//
// El reporte de -race debería desaparecer y el resultado ser siempre -900
// (initialStock - totalSales), en vez de un valor distinto en cada corrida.
func main() {
	stock := initialStock
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(totalSales)
	for i := 0; i < totalSales; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			stock--
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Stock final:", stock) // siempre initialStock - totalSales
}

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

// Hands On 1: reproducir una race condition.
//
// totalSales goroutines decrementan la misma variable stock sin ninguna
// sincronización. Corré este archivo así:
//
//	go run -race race.go
//
// y leé el reporte: Go señala las dos goroutines en conflicto y la línea
// exacta del código.
//
// Nota: a diferencia del snippet de la presentación, acá esperamos con un
// WaitGroup a que todas las goroutines terminen antes de imprimir. El
// WaitGroup no sincroniza el acceso a stock (la race sigue estando en el
// stock--), solo garantiza que todas lleguen a correr antes del Println;
// sin esto el programa podría terminar casi de inmediato sin que -race
// llegue a detectar nada.
func main() {
	stock := initialStock

	var wg sync.WaitGroup
	wg.Add(totalSales)
	for i := 0; i < totalSales; i++ {
		go func() {
			defer wg.Done()
			stock-- // ← data race: totalSales goroutines sin control
		}()
	}
	wg.Wait()

	fmt.Println("Stock final:", stock) // valor impredecible, puede no dar -900
}

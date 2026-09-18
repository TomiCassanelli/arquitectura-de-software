package main

import (
	"fmt"
	"sync"
)

// Hands On 1: reproducir una race condition.
//
// 1000 goroutines decrementan la misma variable stock sin ninguna
// sincronización. Corré este archivo así:
//
//	go run -race race.go
//
// y leé el reporte: Go señala las dos goroutines en conflicto y la línea
// exacta del código.
//
// Nota: a diferencia del snippet de la presentación, acá esperamos con un
// WaitGroup a que las 1000 goroutines terminen antes de imprimir. El
// WaitGroup no sincroniza el acceso a stock (la race sigue estando en el
// stock--), solo garantiza que todas lleguen a correr antes del Println;
// sin esto el programa podría terminar casi de inmediato sin que -race
// llegue a detectar nada.
func main() {
	stock := 100

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			stock-- // ← data race: 1000 goroutines leen y escriben sin control
		}()
	}
	wg.Wait()

	fmt.Println("Stock final:", stock) // valor impredecible, puede no dar -900
}

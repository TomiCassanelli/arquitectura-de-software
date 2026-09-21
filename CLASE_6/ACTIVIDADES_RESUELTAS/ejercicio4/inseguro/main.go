package main

import (
	"fmt"
	"sync"
	"time"
)

// Idéntico a ejercicios-base/ejercicio4/inseguro: se deja acá como punto de
// partida y para poder compararlo lado a lado con mutex/ y channel/. Corré
// go run -race main.go para ver el reporte de data race antes de mirar las
// dos soluciones.
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

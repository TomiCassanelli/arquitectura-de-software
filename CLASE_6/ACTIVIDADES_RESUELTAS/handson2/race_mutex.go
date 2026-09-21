package main

import (
	"fmt"
	"sync"
)

// Hands On 2: corregir la race condition del Hands On 1 con sync.Mutex.
//
// Cambios respecto a handson1/race.go:
//   - Se agrega un sync.Mutex (mu) que protege la sección crítica.
//   - mu.Lock() antes de stock-- y mu.Unlock() después: mientras una
//     goroutine tiene el lock tomado, las otras 999 quedan esperando en
//     Lock() en vez de pisarse entre sí. Esto convierte el
//     "leer-modificar-escribir" en una operación atómica desde el punto de
//     vista de las demás goroutines.
//   - No hace falta ningún otro cambio: el WaitGroup se usa exactamente
//     igual que en el Hands On 1, para esperar a que las 1000 goroutines
//     terminen antes de imprimir.
//
// Verificación: correr de nuevo con go run -race race_mutex.go. El reporte
// de race desaparece y el resultado es siempre -900 (100 - 1000
// decrementos), sin importar cuántas veces se ejecute.
func main() {
	stock := 100
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			stock--
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Stock final:", stock) // siempre -900 (100 - 1000)
}

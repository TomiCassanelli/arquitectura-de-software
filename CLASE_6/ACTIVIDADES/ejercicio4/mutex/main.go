package main

import (
	"fmt"
	"time"
)

// TODO: Ejercicio 4 (a) — Contador seguro con sync.Mutex
//
// Resolver la misma race condition de ejercicio4/inseguro protegiendo el
// contador con sync.Mutex:
//  1. Declarar un sync.Mutex junto al contador de stock.
//  2. Antes de decrementar, mu.Lock(); después, mu.Unlock().
//  3. El time.Sleep que simula la compra puede quedar FUERA del lock (no es
//     parte de la sección crítica).
//  4. Verificar con go run -race main.go que ya no aparece el reporte, y
//     que el resultado es siempre 50.
func main() {
	stock := 100

	// Reemplazar este bloque secuencial por 50 goroutines que decrementan
	// stock protegidas con un sync.Mutex (usar sync.WaitGroup para esperar
	// a que todas terminen antes del Println).
	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond)
		stock--
	}

	fmt.Println("Stock final:", stock)
}

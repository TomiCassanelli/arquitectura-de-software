package main

import (
	"fmt"
	"time"
)

// VARIANTE: el formulario indicará qué bloque copiar y pegar aquí.
const (
	initialStock = XXX
	totalSales   = XXX
	saleDelay    = XXX * time.Millisecond
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
//     que el resultado es initialStock menos totalSales.
func main() {
	stock := initialStock

	// Reemplazar este bloque secuencial por totalSales goroutines que decrementan
	// stock protegidas con un sync.Mutex (usar sync.WaitGroup para esperar
	// a que todas terminen antes del Println).
	for i := 0; i < totalSales; i++ {
		time.Sleep(saleDelay)
		stock--
	}

	fmt.Println("Stock final:", stock)
}

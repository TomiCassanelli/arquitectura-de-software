package main

import (
	"fmt"
	"time"
)

// resultado representa el precio actualizado de un producto.
type resultado struct {
	ID          int
	NuevoPrecio float64
}

// aplicarAumento simula la actualización de precio de un producto contra una
// fuente externa (base de datos, cache, etc). Aplica un 10% de aumento. No
// hace falta tocarla.
func aplicarAumento(id int) resultado {
	time.Sleep(50 * time.Millisecond)
	return resultado{ID: id, NuevoPrecio: float64(id) * 1.10}
}

// procesarSecuencial actualiza todos los productos uno por uno. Sirve como
// base de comparación de tiempos.
func procesarSecuencial(ids []int) []resultado {
	resultados := make([]resultado, 0, len(ids))
	for _, id := range ids {
		resultados = append(resultados, aplicarAumento(id))
	}
	return resultados
}

// TODO: Ejercicio 2 — Worker Pool
//
// Implementar procesarConWorkerPool(ids []int, numWorkers int) []resultado
// que:
//  1. Cree un channel de trabajos (jobs) y encole ahí todos los IDs.
//  2. Lance numWorkers goroutines FIJAS (no una por producto) que tomen IDs
//     de jobs, llamen a aplicarAumento y manden el resultado a un channel de
//     resultados.
//  3. Use sync.WaitGroup para saber cuándo terminaron todos los workers, y
//     así poder cerrar el channel de resultados de forma segura.
//  4. Devuelva todos los resultados juntados en un slice.
//
// Pistas:
//   - El channel de jobs se cierra una sola vez, después de encolar todos
//     los IDs. Los workers salen de su for cuando el channel se cierra y se
//     vacía (range sobre un channel hace esto automáticamente).
//   - El channel de resultados no se puede cerrar desde el mismo lugar
//     donde el main lo está leyendo (range) — pensar en qué otra goroutine
//     podría esperar el WaitGroup y cerrarlo por vos.
func procesarConWorkerPool(ids []int, numWorkers int) []resultado {
	// Reemplazar por la versión con worker pool.
	return procesarSecuencial(ids)
}

func generarIDs(n int) []int {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i + 1
	}
	return ids
}

func main() {
	ids := generarIDs(1000)

	inicio := time.Now()
	procesarSecuencial(ids)
	fmt.Printf("Secuencial:          %v\n", time.Since(inicio))

	for _, n := range []int{5, 20} {
		inicio = time.Now()
		procesarConWorkerPool(ids, n)
		fmt.Printf("Worker pool (N=%2d): %v\n", n, time.Since(inicio))
	}
}

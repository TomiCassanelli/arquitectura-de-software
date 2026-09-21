package main

import (
	"fmt"
	"sync"
	"time"
)

// main es la solución (a) del Ejercicio 4: mismo escenario que
// ejercicio4/inseguro, pero protegiendo el contador con sync.Mutex.
//
// Cambios respecto a la versión insegura:
//   - Se agrega un sync.Mutex (mu) junto al contador.
//   - mu.Lock() / mu.Unlock() rodean únicamente la sección crítica
//     (stock--): mientras una goroutine tiene el lock tomado, las otras 49
//     quedan esperando en Lock() en vez de pisarse entre sí. Esto convierte
//     el "leer, restar, escribir" en una operación atómica desde el punto
//     de vista de las demás goroutines, que es justo lo que evita la race:
//     sin el lock, dos goroutines podían leer el mismo valor viejo antes de
//     que la otra escribiera el nuevo, y una de las dos restas se perdía.
//   - El time.Sleep que simula la compra queda FUERA del lock a propósito:
//     si estuviera adentro, las 50 compras se procesarían de a una,
//     perdiendo el paralelismo real que sí tienen (solo la resta necesita
//     exclusión mutua, no la simulación de latencia).
func main() {
	stock := 100
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)

			mu.Lock()
			stock--
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Stock final:", stock) // siempre 50, y sin reporte de -race
}

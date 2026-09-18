package main

import (
	"fmt"
	"time"
)

// TODO: Ejercicio 4 (b) — Contador seguro con una goroutine dueña del estado
//
// Resolver la misma race condition de ejercicio4/inseguro, pero sin Mutex:
// en vez de proteger memoria compartida, el contador vive DENTRO de una
// única goroutine que es su dueña exclusiva. Las demás goroutines nunca lo
// tocan directamente, solo le piden decrementos por un channel.
//  1. Crear un channel (por ejemplo chan struct{}) para pedir decrementos.
//  2. Lanzar una goroutine "gestora" que arranca con stock := 100 y hace
//     for range sobre ese channel, decrementando stock en cada vuelta.
//  3. Cuando el channel de pedidos se cierra, el range termina; ahí la
//     gestora puede informar el valor final por otro channel.
//  4. Las 50 "compras" solo escriben en el channel de pedidos: nunca leen
//     ni escriben stock directamente.
//
// Pista: esto es lo mismo que le pide el Ejercicio 2 a los workers, pero
// para un contador en vez de una lista de tareas.
func main() {
	stock := 100

	// Reemplazar por la versión basada en channels (una goroutine dueña del
	// estado que recibe pedidos de decremento).
	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond)
		stock--
	}

	fmt.Println("Stock final:", stock)
}

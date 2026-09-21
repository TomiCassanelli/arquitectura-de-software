package main

import (
	"fmt"
	"sync"
	"time"
)

// main es la solución (b) del Ejercicio 4: en vez de proteger memoria
// compartida con un lock, el stock vive dentro de UNA goroutine que es su
// única dueña. Las demás goroutines nunca lo tocan directamente: le piden
// decrementos por un channel.
//
// Cambios respecto a la versión con Mutex:
//   - No existe una variable stock compartida entre goroutines: vive como
//     variable LOCAL dentro de la goroutine "gestora" (el closure que
//     arranca con go func() más abajo).
//   - pedidos es el channel por el que las 50 compras avisan "decrementá
//     stock". Es un chan struct{} (no lleva datos) porque el pedido en sí
//     ya es toda la información que hace falta.
//   - La goroutine gestora hace range sobre pedidos: por cómo funciona un
//     channel en Go, solo una goroutine ejecuta ese cuerpo por vez, así que
//     stock-- nunca corre en paralelo consigo mismo. No hace falta ningún
//     Mutex porque nunca hay dos lecturas/escrituras simultáneas de la
//     misma variable: toda la coordinación pasa por el channel, no por
//     memoria compartida ("no compartas memoria para comunicarte,
//     comunicate para compartir memoria").
//   - Cuando se cierra pedidos (después de wg.Wait() sobre las 50 compras),
//     el range de la gestora termina y recién ahí manda el valor final por
//     el channel resultado, antes de terminar ella misma.
func main() {
	pedidos := make(chan struct{})
	resultado := make(chan int)

	// gestora es la única goroutine que lee y escribe stock.
	go func() {
		stock := 100
		for range pedidos {
			stock--
		}
		resultado <- stock
	}()

	var wg sync.WaitGroup
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			pedidos <- struct{}{} // le pido a la gestora que decremente
		}()
	}

	wg.Wait()
	close(pedidos) // no hay más pedidos: la gestora termina su range y responde

	fmt.Println("Stock final:", <-resultado) // siempre 50, y sin reporte de -race
}

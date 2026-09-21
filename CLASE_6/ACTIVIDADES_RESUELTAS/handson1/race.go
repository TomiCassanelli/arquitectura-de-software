package main

import (
	"fmt"
	"sync"
)

// Hands On 1: reproducir una race condition.
//
// Es el mismo archivo que ejercicios-base/handson1/race.go: el objetivo acá
// no es "resolver" nada, sino observar el bug con go run -race race.go y
// entender qué reporta el detector antes de pasar al Hands On 2.
//
// Qué mirar en el reporte de -race:
//   - Marca DOS accesos en conflicto a la misma dirección de memoria
//     (el stock--), cada uno desde una goroutine distinta.
//   - Al menos uno de los dos accesos es de escritura (por eso es una race:
//     con dos lecturas simultáneas no habría problema).
//   - Te da el stack trace de las dos goroutines y la línea exacta
//     (stock-- dentro del closure), así que se puede ubicar el bug sin
//     tener que razonar todo el programa a mano.
//
// El WaitGroup solo garantiza que las 1000 goroutines lleguen a correr
// antes del Println; NO sincroniza el acceso a stock, por eso la race sigue
// existiendo y el valor final sigue siendo impredecible (compará varias
// corridas: el valor correcto sería -900, pero rara vez da exactamente eso).
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

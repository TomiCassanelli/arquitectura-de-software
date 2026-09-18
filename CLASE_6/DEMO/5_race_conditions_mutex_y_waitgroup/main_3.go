package main

import (
	"fmt"
	"sync"
	"time"
)

var stock = 10
var stockMutex sync.Mutex

func main() {
	// Las goroutines siguen siendo dos; el cambio es proteger stock con Mutex.
	go sellWithMutex("Order A")
	go sellWithMutex("Order B")

	// Solo esperamos para poder ver el resultado.
	time.Sleep(3 * time.Second)
	fmt.Println("Final stock:", stock)
}

func sellWithMutex(order string) {
	// Solo una venta puede pasar esta puerta a la vez.
	stockMutex.Lock()
	// defer libera la llave al terminar esta función.
	defer stockMutex.Unlock()

	fmt.Println(order, "reads stock:", stock)
	time.Sleep(time.Second)
	stock--
	fmt.Println(order, "writes stock:", stock)
}

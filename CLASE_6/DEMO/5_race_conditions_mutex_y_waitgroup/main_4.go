package main

import (
	"fmt"
	"sync"
	"time"
)

var stock = 10
var stockMutex sync.Mutex

func main() {
	// WaitGroup: sirve para que main espere a otras goroutines.
	var sales sync.WaitGroup

	// Ejemplo: vamos a hacer dos ventas, por eso contamos dos tareas pendientes.
	sales.Add(2)

	// &sales significa: pasamos este mismo WaitGroup a cada venta.
	go sellWithMutex("Order A", &sales)
	go sellWithMutex("Order B", &sales)

	// main espera aquí hasta que terminen las dos ventas.
	sales.Wait()
	fmt.Println("Final stock:", stock)
}

func sellWithMutex(order string, sales *sync.WaitGroup) {
	// Done avisa: esta venta terminó.
	defer sales.Done()

	// El Mutex sigue cuidando el dato compartido.
	stockMutex.Lock()
	defer stockMutex.Unlock()

	fmt.Println(order, "reads stock:", stock)
	time.Sleep(time.Second)
	stock--
	fmt.Println(order, "writes stock:", stock)
}

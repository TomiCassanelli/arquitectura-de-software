package main

import (
	"fmt"
	"time"
)

var stock = 10

func main() {
	// Ahora ambas ventas intentan modificar el mismo stock al mismo tiempo.
	go sellWithoutMutex("Order A")
	go sellWithoutMutex("Order B")

	// Solo esperamos para poder ver el resultado.
	time.Sleep(2 * time.Second)
	fmt.Println("Final stock:", stock)
}

func sellWithoutMutex(order string) {
	// Cada goroutine guarda su propia copia del stock que leyó.
	currentStock := stock
	fmt.Println(order, "reads stock:", currentStock)

	// Las dos ventas tienen tiempo de leer 10 antes de que alguna escriba 9.
	time.Sleep(time.Second)

	stock = currentStock - 1
	fmt.Println(order, "writes stock:", stock)
}

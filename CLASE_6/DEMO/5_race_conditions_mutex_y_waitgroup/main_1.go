package main

import "fmt"

func main() {
	stock := 10

	// La venta B recién empieza cuando la venta A ya terminó.
	stock = sell(stock, "Order A")
	stock = sell(stock, "Order B")

	fmt.Println("Final stock:", stock)
}

func sell(stock int, order string) int {
	// Esta venta recibe el stock más actualizado porque todo es secuencial.
	fmt.Println(order, "reads stock:", stock)
	stock--
	fmt.Println(order, "writes stock:", stock)
	return stock
}

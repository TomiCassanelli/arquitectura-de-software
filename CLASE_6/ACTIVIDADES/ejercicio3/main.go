package main

import (
	"fmt"
	"time"
)

type Producto struct {
	ID      int
	Precio  float64
	Stock   int
	Reviews float64 // 0 si no llegaron a tiempo
}

func consultarPrecio(id int) float64 {
	time.Sleep(200 * time.Millisecond)
	return 199.99
}

func consultarStock(id int) int {
	time.Sleep(150 * time.Millisecond)
	return 42
}

// consultarReviews simula una fuente deliberadamente lenta (2s) para forzar
// el timeout que hay que implementar en este ejercicio.
func consultarReviews(id int) float64 {
	time.Sleep(2 * time.Second)
	return 4.5
}

// obtenerFichaConcurrente es la solución del Ejercicio 1: fan-out/fan-in sin
// ningún límite de tiempo. Con la consultarReviews de este archivo, reviews
// bloquea toda la respuesta 2 segundos enteros.
func obtenerFichaConcurrente(id int) Producto {
	precioCh := make(chan float64, 1)
	stockCh := make(chan int, 1)
	reviewsCh := make(chan float64, 1)

	go func() { precioCh <- consultarPrecio(id) }()
	go func() { stockCh <- consultarStock(id) }()
	go func() { reviewsCh <- consultarReviews(id) }()

	return Producto{
		ID:      id,
		Precio:  <-precioCh,
		Stock:   <-stockCh,
		Reviews: <-reviewsCh,
	}
}

// TODO: Ejercicio 3 — Context: cortar una fuente lenta
//
// Implementar obtenerFichaConTimeout(id int) Producto que:
//  1. Cree un context.WithTimeout de 300ms (con su defer cancel()).
//  2. Lance precio, stock y reviews igual que en obtenerFichaConcurrente.
//  3. Al esperar el resultado de reviews, use select entre reviewsCh y
//     ctx.Done(): si el contexto se cancela primero, la ficha se devuelve
//     con Reviews en 0 (o el default que prefieras) en vez de bloquear.
//
// Pista: precio y stock no necesitan timeout en este ejercicio (son
// rápidos); alcanza con protegerse contra reviews.
func obtenerFichaConTimeout(id int) Producto {
	// Reemplazar por la versión con context.WithTimeout + select.
	return obtenerFichaConcurrente(id)
}

func main() {
	inicio := time.Now()
	ficha := obtenerFichaConTimeout(1)
	fmt.Printf("Ficha: %+v (tardó %v)\n", ficha, time.Since(inicio))
}

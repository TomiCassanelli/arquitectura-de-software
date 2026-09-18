package main

import (
	"fmt"
	"time"
)

// Producto es la ficha completa que se devuelve al cliente.
type Producto struct {
	ID      int
	Precio  float64
	Stock   int
	Reviews float64
}

// Las siguientes tres funciones simulan llamadas a fuentes externas (base de
// datos, servicio de stock, API de reviews) con latencias distintas. No hace
// falta tocarlas.
func consultarPrecio(id int) float64 {
	time.Sleep(200 * time.Millisecond)
	return 199.99
}

func consultarStock(id int) int {
	time.Sleep(150 * time.Millisecond)
	return 42
}

func consultarReviews(id int) float64 {
	time.Sleep(300 * time.Millisecond)
	return 4.5
}

// obtenerFichaSecuencial arma la ficha llamando a las tres fuentes una
// después de la otra. Sirve como base de comparación de tiempos.
func obtenerFichaSecuencial(id int) Producto {
	precio := consultarPrecio(id)
	stock := consultarStock(id)
	reviews := consultarReviews(id)
	return Producto{ID: id, Precio: precio, Stock: stock, Reviews: reviews}
}

// TODO: Ejercicio 1 — Fan-Out/Fan-In
//
// Implementar obtenerFichaConcurrente(id int) Producto que:
//  1. Lance las tres consultas como goroutines independientes (fan-out).
//  2. Recolecte los tres resultados (fan-in), por ejemplo con un channel por
//     cada fuente, o con un channel compartido + sync.WaitGroup.
//  3. Devuelva la Producto ya armada con los tres valores.
//
// Pista: un channel con buffer 1 por cada fuente es la forma más simple de
// empezar (la goroutine que lo llena nunca se bloquea al escribir).
func obtenerFichaConcurrente(id int) Producto {
	// Reemplazar esta implementación de ejemplo (secuencial) por la versión
	// concurrente.
	return obtenerFichaSecuencial(id)
}

func main() {
	inicio := time.Now()
	fichaSecuencial := obtenerFichaSecuencial(1)
	fmt.Printf("Secuencial:  %+v (tardó %v)\n", fichaSecuencial, time.Since(inicio))

	inicio = time.Now()
	fichaConcurrente := obtenerFichaConcurrente(1)
	fmt.Printf("Concurrente: %+v (tardó %v)\n", fichaConcurrente, time.Since(inicio))
}

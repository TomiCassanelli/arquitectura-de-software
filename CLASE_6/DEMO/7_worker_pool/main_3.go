package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	productIDs := []int{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113}

	// Por este channel llegan los productos que hay que actualizar.
	productChannel := make(chan int)
	var finishedWorkers sync.WaitGroup

	// Abrimos solo tres trabajadores, aunque haya cinco productos.
	for workerID := 1; workerID <= 3; workerID++ {
		finishedWorkers.Add(1)
		go updateProducts(workerID, productChannel, &finishedWorkers)
	}

	// Como no tiene buffer, main entrega cada producto al próximo worker libre.
	// No hay un reparto fijo: puede cambiar qué worker toma cada producto.
	for _, productID := range productIDs {
		productChannel <- productID
	}

	// Ya no llegarán más productos: los trabajadores pueden terminar.
	close(productChannel)

	// main espera a que terminen los tres trabajadores.
	finishedWorkers.Wait()
	fmt.Println("All products updated")
}

func updateProducts(workerID int, productChannel chan int, finishedWorkers *sync.WaitGroup) {
	defer finishedWorkers.Done()

	// Cuando termina un producto, este worker vuelve a esperar otro.
	// range termina cuando main cerró el channel y ya no quedan productos.
	for productID := range productChannel {
		fmt.Printf("Worker %d updates product %d\n", workerID, productID)
		time.Sleep(1000 * time.Millisecond)
	}
}

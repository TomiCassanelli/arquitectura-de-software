package main

import (
	"fmt"
	"time"
)

func main() {
	productIDs := []int{101, 102, 103, 104, 105}

	// Primero se actualiza un producto y luego empieza el siguiente.
	for _, productID := range productIDs {
		updateProductPrice(productID)
	}

	fmt.Println("All products updated")
}

func updateProductPrice(productID int) {
	fmt.Println("Updating product", productID)
	time.Sleep(500 * time.Millisecond)
}

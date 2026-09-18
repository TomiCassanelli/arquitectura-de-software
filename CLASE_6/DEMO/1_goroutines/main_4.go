package main

import (
	"fmt"
	"time"
)

func main() {
	go processUsers()

	// Hay tiempo para que terminen los tres usuarios.
	time.Sleep(2 * time.Second)
	fmt.Println("Main finished")
}

func processUsers() {
	const totalUsers = 3

	for userID := 1; userID <= totalUsers; userID++ {
		// Este go hace que cada usuario avance de forma independiente.
		go processUser(userID)
	}
}

func processUser(userID int) {
	fmt.Println("Start user", userID)
	time.Sleep(time.Second)
	fmt.Println("End user", userID)
}

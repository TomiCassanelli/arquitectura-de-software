package main

import (
	"fmt"
	"time"
)

func main() {
	go processUsers()

	// Ahora main no termina enseguida: deja trabajar a la goroutine.
	time.Sleep(4 * time.Second)
	fmt.Println("Main finished")
}

func processUsers() {
	const totalUsers = 3

	for userID := 1; userID <= totalUsers; userID++ {
		processUser(userID)
	}
}

func processUser(userID int) {
	fmt.Println("Start user", userID)
	time.Sleep(time.Second)
	fmt.Println("End user", userID)
}

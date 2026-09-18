package main

import (
	"fmt"
	"time"
)

func main() {
	// Solo agregamos go: processUsers empieza por separado de main.
	go processUsers()

	// main termina enseguida y el programa se cierra.
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

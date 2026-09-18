package main

import (
	"fmt"
	"time"
)

func main() {
	// Todavía no hay goroutines: main hace todo el trabajo.
	processUsers()
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

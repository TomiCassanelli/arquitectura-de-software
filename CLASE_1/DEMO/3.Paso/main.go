// Sabiendo cuál es el contrato, creamos la estructura que viaja por la red. Esta es la que se usará en producción.

package main

import "fmt"

// --- LO DE LOS PASOS ANTERIORES ---
type User struct {
	ID     int64
	Name   string
	Active bool
}

type UserRepo interface {
	GetUserByID(id int64) (User, error)
}

// ==============================================================================
// NUEVO EN ESTE PASO: IMPLEMENTACIÓN REAL (clients)
// En el TP real, esto iría en la carpeta /clients.
// Esta es la estructura que viajará por la red.
// ==============================================================================

// UserHTTP es el adaptador real.
type UserHTTP struct{}

// Es la forma de decirle a Go: "Esta función le pertenece al struct UserHTTP". Al hacer esto, UserHTTP cumple automáticamente con el contrato de la interfaz UserRepo que armamos antes, porque ahora posee el método GetUserByID.
func (repo UserHTTP) GetUserByID(id int64) (User, error) {
	fmt.Println("🌐 [RED] Viajando por internet hacia el Microservicio de Usuarios...")

	// Nota: Simulamos la respuesta final devolviendo la estructura armada:
	return User{ID: id, Name: "Socio Real de BD", Active: true}, nil
}

func main() {
	fmt.Println("Paso 3: Ya tenemos nuestro cliente HTTP para ir a producción.")
}

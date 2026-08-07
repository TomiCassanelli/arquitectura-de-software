// Ahora sumamos la implementación falsa. Este es el salvavidas para cuando no hay internet o el otro microservicio no está listo.

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

type UserHTTP struct{}

func (repo UserHTTP) GetUserByID(id int64) (User, error) {
	fmt.Println("🌐 [RED] Viajando por internet hacia el Microservicio de Usuarios...")
	return User{ID: id, Name: "Socio Real de BD", Active: true}, nil
}

// ==============================================================================
// NUEVO EN ESTE PASO: EL MOCK PARA DESARROLLO LOCAL Y TESTING (clients)
// ==============================================================================

// UserMock es el adaptador falso.
type UserMock struct{}

// Cumple con el contrato, pero responde al instante sin usar internet.
func (repo UserMock) GetUserByID(id int64) (User, error) {
	fmt.Println("[MOCK] Devolviendo datos simulados al instante (Sin usar red)...")
	return User{ID: id, Name: "Socio de Prueba (Mock)", Active: true}, nil
}

func main() {
	fmt.Println("Paso 4: Ya tenemos el Mock creado para poder testear sin conexión.")
}

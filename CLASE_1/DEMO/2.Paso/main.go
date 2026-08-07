// Agregamos la interfaz. Les explicamos que antes de programar cómo vamos a buscar los datos, primero definimos qué necesitamos (el contrato).

package main

import "fmt"

// --- LO DEL PASO ANTERIOR ---
type User struct {
	ID     int64
	Name   string
	Active bool
}

// ==============================================================================
// NUEVO EN ESTE PASO: LA CAPA DE CLIENTES - EL CONTRATO (clients)
// En el TP real, esto iría en la carpeta /clients.
// Es la capa donde se consume servicios externos: APIs de terceros u otros microservicios, por ejemplo.
// ==============================================================================

// UserRepo es nuestro "contrato". Le dice al sistema: "Cualquier adaptador que
// uses para traer usuarios DEBE tener obligatoriamente el método GetUserByID".
type UserRepo interface {
	GetUserByID(id int64) (User, error)
}

func main() {
	fmt.Println("Paso 2: Entidad User y Contrato UserRepo creados.")
}

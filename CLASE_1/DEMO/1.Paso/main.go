// En este primer paso, solo creamos la estructura de datos que representa nuestro negocio.

package main

import "fmt"

// ==============================================================================
// NUEVO EN ESTE PASO: LA CAPA DE DOMINIO (domain)
// En el TP real, esto iría en la carpeta /domain
// Define las entidades centrales del negocio sin depender de nada externo.
// ==============================================================================

// User representa a un socio del gimnasio.
type User struct {
	ID     int64
	Name   string
	Active bool
}

func main() {
	fmt.Println("Paso 1: Entidad User creada.")
}

// Introducimos la lógica de la petición web. El controlador delega la búsqueda de datos en la interfaz, sin importarle qué implementación se use.

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// --- LO DE LOS PASOS ANTERIORES (Entidad, Interfaz, HTTP y Mock) ---
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
	return User{ID: id, Name: "Socio Real", Active: true}, nil
}

type UserMock struct{}

func (repo UserMock) GetUserByID(id int64) (User, error) {
	fmt.Println("🧪 [MOCK] Devolviendo datos simulados...")
	return User{ID: id, Name: "Socio de Prueba (Mock)", Active: true}, nil
}

// ==============================================================================
// NUEVO EN ESTE PASO: LA CAPA DE CONTROLADORES (controllers)
// En el TP real, esto iría en la carpeta /controllers.
// ==============================================================================

// ActivitiesController maneja las reservas de clases.
type ActivitiesController struct {
	// Fíjense que NO decimos "UserClient UserHTTP" ni "UserClient UserMock".
	// Le pasamos la INTERFAZ "UserRepo" (El contrato).

	UserClient UserRepo
}

// Reserve proceas la reserva de un turno.
func (c *ActivitiesController) Reserve(ctx *gin.Context) {
	socioID := int64(123)

	// Llamamos al método del contrato, no sabe aún, el confía en la respuesta.
	user, err := c.UserClient.GetUserByID(socioID)
	if err != nil || !user.Active {
		ctx.String(400, "Error: No se pudo verificar al socio o no está activo")
		return
	}

	mensaje := fmt.Sprintf("¡Éxito! Reserva confirmada para: %s", user.Name)
	fmt.Println(mensaje)
	ctx.String(200, mensaje)
}

func main() {
	fmt.Println("Paso 5: Controlador creado y dependiendo exclusivamente de la interfaz.")
}

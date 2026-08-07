// El último paso une todas las piezas. En la función main inyectamos la dependencia (el Mock) al controlador y levantamos el servidor para recibir peticiones reales.

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID     int64
	Name   string
	Active bool
}

type UserRepo interface{ GetUserByID(id int64) (User, error) }

type UserHTTP struct{}

func (repo UserHTTP) GetUserByID(id int64) (User, error) {
	fmt.Println("🌐 [RED] Viajando por internet hacia el Microservicio de Usuarios...")
	return User{ID: id, Name: "Socio Real", Active: true}, nil
}

type UserMock struct{}

func (repo UserMock) GetUserByID(id int64) (User, error) {
	fmt.Println("🧪 [MOCK] Devolviendo datos simulados...")
	return User{ID: id, Name: "Socio de Prueba (Mock)", Active: true}, nil
}

type ActivitiesController struct {
	UserClient UserRepo
}

func (c *ActivitiesController) Reserve(ctx *gin.Context) {
	socioID := int64(123)
	user, err := c.UserClient.GetUserByID(socioID)
	if err != nil || !user.Active {
		ctx.String(400, "Error: No se pudo verificar al socio o no está activo")
		return
	}

	mensaje := fmt.Sprintf("¡Éxito! Reserva confirmada para: %s", user.Name)
	fmt.Println(mensaje)
	ctx.String(200, mensaje)
}

// ==============================================================================
// NUEVO EN ESTE PASO: EL PUNTO DE ENTRADA Y SERVIDOR (main.go)
// ==============================================================================

func main() {

	// INYECCIÓN DE DEPENDENCIAS:
	// Le pasamos el Mock. Para ir a producción, solo se cambia UserMock{} por UserHTTP{}.
	actCtrl := &ActivitiesController{
		UserClient: UserMock{},
	}

	router := gin.Default()
	router.POST("/activities/reserve", func(ctx *gin.Context) {
		actCtrl.Reserve(ctx)
	})

	fmt.Println("🚀 Microservicio de Actividades levantado en el puerto 8082...")
	fmt.Println("👉 Probá hacer un POST (o abrilo en tu navegador si cambiás el POST por GET) a http://localhost:8082/activities/reserve")

	if err := router.Run(":8082"); err != nil {
		fmt.Printf("Error fatal: %s\n", err)
	}
}

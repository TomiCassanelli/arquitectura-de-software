package main

import (
	"fmt"

	"practico/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)

	r.GET("/activities", controllers.GetAllActivities)
	r.POST("/activities/:id/reserve", controllers.ReserveActivity)

	r.GET("/products", controllers.GetCatalog)
	r.POST("/products/buy", controllers.BuyProduct)

	r.GET("/users/:id/points", controllers.GetPoints)
	r.POST("/rewards/redeem", controllers.RedeemReward)

	fmt.Println("Iniciando servidor Gin en el puerto 8080...")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}

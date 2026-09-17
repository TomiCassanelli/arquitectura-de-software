package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// En el monolito, clientes, productos y pedidos viven en el mismo proceso.
	router.POST("/clientes", crearCliente)
	router.GET("/clientes/:id", obtenerCliente)
	router.GET("/productos", listarProductos)
	router.POST("/pedidos", confirmarPedido)

	log.Println("Monolito escuchando en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

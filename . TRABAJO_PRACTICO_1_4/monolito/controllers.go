package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func crearCliente(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"mensaje": "cliente creado"})
}

func obtenerCliente(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":     c.Param("id"),
		"nombre": "Ana Pérez",
	})
}

func listarProductos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"productos": []gin.H{
			{"id": "P-1", "nombre": "Auriculares", "stock": 10},
			{"id": "P-2", "nombre": "Teclado", "stock": 8},
		},
	})
}

func confirmarPedido(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"mensaje":   "pedido confirmado",
		"pedido_id": "PED-1",
	})
}

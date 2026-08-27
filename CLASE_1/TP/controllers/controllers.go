package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Usuarios
func CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "Usuario creado exitosamente")
}

func GetUserByID(c *gin.Context) {
	// id := c.Param("id") // si se necesita
	c.String(http.StatusOK, "Detalle del usuario y estado de cuota")
}

// Actividades
func GetAllActivities(c *gin.Context) {
	c.String(http.StatusOK, "Grilla de horarios de Hyrox y Yoga")
}

func ReserveActivity(c *gin.Context) {
	c.String(http.StatusOK, "Reserva de cupo confirmada")
}

// Tienda / Café
func GetCatalog(c *gin.Context) {
	c.String(http.StatusOK, "Catálogo de suplementos, ropa y menú del café")
}

func BuyProduct(c *gin.Context) {
	c.String(http.StatusOK, "Compra procesada exitosamente")
}

// Beneficios / Loyalty
func GetPoints(c *gin.Context) {
	c.String(http.StatusOK, "Puntos acumulados del socio")
}

func RedeemReward(c *gin.Context) {
	c.String(http.StatusOK, "Beneficio canjeado en comercio adherido")
}

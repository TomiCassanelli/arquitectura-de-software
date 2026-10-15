package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CompraOK simula tres pasos de una compra exitosa con el mismo track_id.
func (d *Demo) CompraOK(c *gin.Context) {
	id := c.GetString("track_id")
	d.Logs.WriteTrack(id, "info", "compra recibida")
	d.Logs.WriteTrack(id, "info", "stock verificado")
	d.Logs.WriteTrack(id, "info", "pago aprobado")
	c.JSON(http.StatusOK, gin.H{"resultado": "compra confirmada", "track_id": id})
}

// CompraError simula el mismo recorrido hasta que falla el pago.
func (d *Demo) CompraError(c *gin.Context) {
	id := c.GetString("track_id")
	d.Logs.WriteTrack(id, "info", "compra recibida")
	d.Logs.WriteTrack(id, "info", "stock verificado")
	d.Logs.WriteTrack(id, "error", "pago rechazado")
	c.JSON(http.StatusInternalServerError, gin.H{"error": "pago rechazado", "track_id": id})
}
